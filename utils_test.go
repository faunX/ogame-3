package ogame

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCoord(t *testing.T) {
	coord, _ := ParseCoord("[P:1:2:3]")
	assert.Equal(t, Coordinate{1, 2, 3, PlanetType}, coord)
	coord, _ = ParseCoord("[M:1:2:3]")
	assert.Equal(t, Coordinate{1, 2, 3, MoonType}, coord)
	coord, _ = ParseCoord("M:1:2:3")
	assert.Equal(t, Coordinate{1, 2, 3, MoonType}, coord)
	coord, _ = ParseCoord("1:2:3")
	assert.Equal(t, Coordinate{1, 2, 3, PlanetType}, coord)
	coord, _ = ParseCoord("1:2:3")
	assert.Equal(t, Coordinate{1, 2, 3, PlanetType}, coord)
	coord, _ = ParseCoord("D:1:2:3")
	assert.Equal(t, Coordinate{1, 2, 3, DebrisType}, coord)
	coord, _ = ParseCoord("[D:1:2:3]")
	assert.Equal(t, Coordinate{1, 2, 3, DebrisType}, coord)
	_, err := ParseCoord("[A:1:2:3]")
	assert.NotNil(t, err)
	_, err = ParseCoord("aP:1:2:3")
	assert.NotNil(t, err)
	_, err = ParseCoord("P:1234:2:3")
	assert.NotNil(t, err)
	_, err = ParseCoord("P:1:2345:3")
	assert.NotNil(t, err)
	_, err = ParseCoord("P:1:2:3456")
	assert.NotNil(t, err)
}

func TestName2id(t *testing.T) {
	assert.Equal(t, ID(0), DefenceName2ID("Not valid"))
	assert.Equal(t, RocketLauncherID, DefenceName2ID("Rocket Launcher"))
	assert.Equal(t, LightLaserID, DefenceName2ID("Light Laser"))
	assert.Equal(t, HeavyLaserID, DefenceName2ID("Heavy Laser"))
	assert.Equal(t, GaussCannonID, DefenceName2ID("Gauss Cannon"))
	assert.Equal(t, IonCannonID, DefenceName2ID("Ion Cannon"))
	assert.Equal(t, PlasmaTurretID, DefenceName2ID("Plasma Turret"))
	assert.Equal(t, SmallShieldDomeID, DefenceName2ID("Small Shield Dome"))
	assert.Equal(t, LargeShieldDomeID, DefenceName2ID("Large Shield Dome"))
	assert.Equal(t, AntiBallisticMissilesID, DefenceName2ID("Anti-Ballistic Missiles"))
	assert.Equal(t, InterplanetaryMissilesID, DefenceName2ID("Interplanetary Missiles"))

	// hu
	assert.Equal(t, RocketLauncherID, DefenceName2ID("Rakéta kilövő"))
	assert.Equal(t, LightLaserID, DefenceName2ID("Könnyű lézer"))
	assert.Equal(t, HeavyLaserID, DefenceName2ID("Nehéz lézer"))
	assert.Equal(t, GaussCannonID, DefenceName2ID("Gauss ágyú"))
	assert.Equal(t, IonCannonID, DefenceName2ID("Ion ágyú"))
	assert.Equal(t, PlasmaTurretID, DefenceName2ID("Plazmatorony"))
	assert.Equal(t, SmallShieldDomeID, DefenceName2ID("Kis pajzskupola"))
	assert.Equal(t, LargeShieldDomeID, DefenceName2ID("Nagy pajzskupola"))
	assert.Equal(t, AntiBallisticMissilesID, DefenceName2ID("Anti-Ballasztikus rakéták"))
	assert.Equal(t, InterplanetaryMissilesID, DefenceName2ID("Bolygóközi rakéták"))

	// si
	assert.Equal(t, RocketLauncherID, DefenceName2ID("Raketnik"))
	assert.Equal(t, LightLaserID, DefenceName2ID("Lahki laser"))
	assert.Equal(t, HeavyLaserID, DefenceName2ID("Težek laser"))
	assert.Equal(t, GaussCannonID, DefenceName2ID("Gaussov top"))
	assert.Equal(t, IonCannonID, DefenceName2ID("Ionski top"))
	assert.Equal(t, PlasmaTurretID, DefenceName2ID("Plazemski top"))
	assert.Equal(t, SmallShieldDomeID, DefenceName2ID("Majhen ščit"))
	assert.Equal(t, LargeShieldDomeID, DefenceName2ID("Velik ščit"))
	assert.Equal(t, AntiBallisticMissilesID, DefenceName2ID("Protibalistične rakete"))
	assert.Equal(t, InterplanetaryMissilesID, DefenceName2ID("Medplanetarne rakete"))

	// ro
	assert.Equal(t, RocketLauncherID, DefenceName2ID("Lansator de Rachete"))
	assert.Equal(t, LightLaserID, DefenceName2ID("Laser Usor"))
	assert.Equal(t, HeavyLaserID, DefenceName2ID("Laser Greu"))
	assert.Equal(t, GaussCannonID, DefenceName2ID("Tun Gauss"))
	assert.Equal(t, IonCannonID, DefenceName2ID("Tun Magnetic"))
	assert.Equal(t, PlasmaTurretID, DefenceName2ID("Turela cu Plasma"))
	assert.Equal(t, SmallShieldDomeID, DefenceName2ID("Scut Planetar Mic"))
	assert.Equal(t, LargeShieldDomeID, DefenceName2ID("Scut Planetar Mare"))
	assert.Equal(t, AntiBallisticMissilesID, DefenceName2ID("Racheta Anti-Balistica"))
	assert.Equal(t, InterplanetaryMissilesID, DefenceName2ID("Rachete Interplanetare"))

	// sk
	assert.Equal(t, RocketLauncherID, DefenceName2ID("Raketový komplet"))
	assert.Equal(t, LightLaserID, DefenceName2ID("Ľahký laser"))
	assert.Equal(t, HeavyLaserID, DefenceName2ID("Ťažký laser"))
	assert.Equal(t, GaussCannonID, DefenceName2ID("Gaussov kanón"))
	assert.Equal(t, IonCannonID, DefenceName2ID("Iónový kanón"))
	assert.Equal(t, PlasmaTurretID, DefenceName2ID("Plazmová veža"))
	assert.Equal(t, SmallShieldDomeID, DefenceName2ID("Malý planetárny štít"))
	assert.Equal(t, LargeShieldDomeID, DefenceName2ID("Veľký planetárny štít"))
	assert.Equal(t, AntiBallisticMissilesID, DefenceName2ID("Protiraketové strely"))
	assert.Equal(t, InterplanetaryMissilesID, DefenceName2ID("Medziplanetárne rakety"))

	// gr
	assert.Equal(t, RocketLauncherID, DefenceName2ID("Εκτοξευτής Πυραύλων"))
	assert.Equal(t, LightLaserID, DefenceName2ID("Ελαφρύ Λέιζερ"))
	assert.Equal(t, HeavyLaserID, DefenceName2ID("Βαρύ Λέιζερ"))
	assert.Equal(t, GaussCannonID, DefenceName2ID("Κανόνι Gauss"))
	assert.Equal(t, IonCannonID, DefenceName2ID("Κανόνι Ιόντων"))
	assert.Equal(t, PlasmaTurretID, DefenceName2ID("Πυργίσκοι Πλάσματος"))
	assert.Equal(t, SmallShieldDomeID, DefenceName2ID("Μικρός Αμυντικός Θόλος"))
	assert.Equal(t, LargeShieldDomeID, DefenceName2ID("Μεγάλος Αμυντικός Θόλος"))
	assert.Equal(t, AntiBallisticMissilesID, DefenceName2ID("Αντι-Βαλλιστικοί Πύραυλοι"))
	assert.Equal(t, InterplanetaryMissilesID, DefenceName2ID("Διαπλανητικοί Πύραυλοι"))

	// tw
	assert.Equal(t, RocketLauncherID, DefenceName2ID("飛彈發射器"))
	assert.Equal(t, LightLaserID, DefenceName2ID("輕型雷射炮"))
	assert.Equal(t, HeavyLaserID, DefenceName2ID("重型雷射炮"))
	assert.Equal(t, GaussCannonID, DefenceName2ID("磁軌炮"))
	assert.Equal(t, IonCannonID, DefenceName2ID("離子加農炮"))
	assert.Equal(t, PlasmaTurretID, DefenceName2ID("電漿炮塔"))
	assert.Equal(t, SmallShieldDomeID, DefenceName2ID("小型防護罩"))
	assert.Equal(t, LargeShieldDomeID, DefenceName2ID("大型防護罩"))
	assert.Equal(t, AntiBallisticMissilesID, DefenceName2ID("反彈道導彈"))
	assert.Equal(t, InterplanetaryMissilesID, DefenceName2ID("星際導彈"))

	// hr
	assert.Equal(t, RocketLauncherID, DefenceName2ID("Raketobacači"))
	assert.Equal(t, LightLaserID, DefenceName2ID("Mali laser"))
	assert.Equal(t, HeavyLaserID, DefenceName2ID("Veliki laser"))
	assert.Equal(t, GaussCannonID, DefenceName2ID("Gaussov top"))
	assert.Equal(t, IonCannonID, DefenceName2ID("Ionski top"))
	assert.Equal(t, PlasmaTurretID, DefenceName2ID("Plazma top"))
	assert.Equal(t, SmallShieldDomeID, DefenceName2ID("Mala štitna kupola"))
	assert.Equal(t, LargeShieldDomeID, DefenceName2ID("Velika štitna kupola"))
	assert.Equal(t, AntiBallisticMissilesID, DefenceName2ID("Anti-balističke rakete"))
	assert.Equal(t, InterplanetaryMissilesID, DefenceName2ID("Interplanetarne rakete"))

	// mx
	assert.Equal(t, RocketLauncherID, DefenceName2ID("Lanzamisiles"))
	assert.Equal(t, LightLaserID, DefenceName2ID("Láser pequeño"))
	assert.Equal(t, HeavyLaserID, DefenceName2ID("Láser grande"))
	assert.Equal(t, GaussCannonID, DefenceName2ID("Cañón Gauss"))
	assert.Equal(t, IonCannonID, DefenceName2ID("Cañón iónico"))
	assert.Equal(t, PlasmaTurretID, DefenceName2ID("Cañón de plasma"))
	assert.Equal(t, SmallShieldDomeID, DefenceName2ID("Cúpula pequeña de protección"))
	assert.Equal(t, LargeShieldDomeID, DefenceName2ID("Cúpula grande de protección"))
	assert.Equal(t, AntiBallisticMissilesID, DefenceName2ID("Misil de intercepción"))
	assert.Equal(t, InterplanetaryMissilesID, DefenceName2ID("Misil interplanetario"))

	// cz
	assert.Equal(t, RocketLauncherID, DefenceName2ID("Raketomet"))
	assert.Equal(t, LightLaserID, DefenceName2ID("Lehký laser"))
	assert.Equal(t, HeavyLaserID, DefenceName2ID("Těžký laser"))
	assert.Equal(t, GaussCannonID, DefenceName2ID("Gaussův kanón"))
	assert.Equal(t, IonCannonID, DefenceName2ID("Iontový kanón"))
	assert.Equal(t, PlasmaTurretID, DefenceName2ID("Plasmová věž"))
	assert.Equal(t, SmallShieldDomeID, DefenceName2ID("Malý planetární štít"))
	assert.Equal(t, LargeShieldDomeID, DefenceName2ID("Velký planetární štít"))
	assert.Equal(t, AntiBallisticMissilesID, DefenceName2ID("Antibalistické rakety"))
	assert.Equal(t, InterplanetaryMissilesID, DefenceName2ID("Meziplanetární rakety"))

	// it
	assert.Equal(t, RocketLauncherID, DefenceName2ID("Lanciamissili"))
	assert.Equal(t, LightLaserID, DefenceName2ID("Laser leggero"))
	assert.Equal(t, HeavyLaserID, DefenceName2ID("Laser pesante"))
	assert.Equal(t, GaussCannonID, DefenceName2ID("Cannone Gauss"))
	assert.Equal(t, IonCannonID, DefenceName2ID("Cannone ionico"))
	assert.Equal(t, PlasmaTurretID, DefenceName2ID("Cannone al Plasma"))
	assert.Equal(t, SmallShieldDomeID, DefenceName2ID("Cupola scudo piccola"))
	assert.Equal(t, LargeShieldDomeID, DefenceName2ID("Cupola scudo potenziata"))
	assert.Equal(t, AntiBallisticMissilesID, DefenceName2ID("Missili anti balistici"))
	assert.Equal(t, InterplanetaryMissilesID, DefenceName2ID("Missili Interplanetari"))

	// de
	assert.Equal(t, RocketLauncherID, DefenceName2ID("Raketenwerfer"))
	assert.Equal(t, LightLaserID, DefenceName2ID("Leichtes Lasergeschütz"))
	assert.Equal(t, HeavyLaserID, DefenceName2ID("Schweres Lasergeschütz"))
	assert.Equal(t, GaussCannonID, DefenceName2ID("Gaußkanone"))
	assert.Equal(t, IonCannonID, DefenceName2ID("Ionengeschütz"))
	assert.Equal(t, PlasmaTurretID, DefenceName2ID("Plasmawerfer"))
	assert.Equal(t, SmallShieldDomeID, DefenceName2ID("Kleine Schildkuppel"))
	assert.Equal(t, LargeShieldDomeID, DefenceName2ID("Große Schildkuppel"))
	assert.Equal(t, AntiBallisticMissilesID, DefenceName2ID("Abfangrakete"))
	assert.Equal(t, InterplanetaryMissilesID, DefenceName2ID("Interplanetarrakete"))

	// dk
	assert.Equal(t, RocketLauncherID, DefenceName2ID("Raketkanon"))
	assert.Equal(t, LightLaserID, DefenceName2ID("Lille Laserkanon"))
	assert.Equal(t, HeavyLaserID, DefenceName2ID("Stor Laserkanon"))
	assert.Equal(t, GaussCannonID, DefenceName2ID("Gausskanon"))
	assert.Equal(t, IonCannonID, DefenceName2ID("Ionkanon"))
	assert.Equal(t, PlasmaTurretID, DefenceName2ID("Plasmakanon"))
	assert.Equal(t, SmallShieldDomeID, DefenceName2ID("Lille Planetskjold"))
	assert.Equal(t, LargeShieldDomeID, DefenceName2ID("Stort Planetskjold"))
	assert.Equal(t, AntiBallisticMissilesID, DefenceName2ID("Forsvarsraket"))
	assert.Equal(t, InterplanetaryMissilesID, DefenceName2ID("Interplanetarraket"))

	// es
	assert.Equal(t, RocketLauncherID, DefenceName2ID("Lanzamisiles"))
	assert.Equal(t, LightLaserID, DefenceName2ID("Láser pequeño"))
	assert.Equal(t, HeavyLaserID, DefenceName2ID("Láser grande"))
	assert.Equal(t, GaussCannonID, DefenceName2ID("Cañón gauss"))
	assert.Equal(t, IonCannonID, DefenceName2ID("Cañón iónico"))
	assert.Equal(t, PlasmaTurretID, DefenceName2ID("Cañón de plasma"))
	assert.Equal(t, SmallShieldDomeID, DefenceName2ID("Cúpula pequeña de protección"))
	assert.Equal(t, LargeShieldDomeID, DefenceName2ID("Cúpula grande de protección"))
	assert.Equal(t, AntiBallisticMissilesID, DefenceName2ID("Misiles antibalísticos"))
	assert.Equal(t, InterplanetaryMissilesID, DefenceName2ID("Misil interplanetario"))

	// fr
	assert.Equal(t, RocketLauncherID, DefenceName2ID("Lanceur de missiles"))
	assert.Equal(t, LightLaserID, DefenceName2ID("Artillerie laser légère"))
	assert.Equal(t, HeavyLaserID, DefenceName2ID("Artillerie laser lourde"))
	assert.Equal(t, GaussCannonID, DefenceName2ID("Canon de Gauss"))
	assert.Equal(t, IonCannonID, DefenceName2ID("Artillerie à ions"))
	assert.Equal(t, PlasmaTurretID, DefenceName2ID("Lanceur de plasma"))
	assert.Equal(t, SmallShieldDomeID, DefenceName2ID("Petit bouclier"))
	assert.Equal(t, LargeShieldDomeID, DefenceName2ID("Grand bouclier"))
	assert.Equal(t, AntiBallisticMissilesID, DefenceName2ID("Missile d`interception"))
	assert.Equal(t, InterplanetaryMissilesID, DefenceName2ID("Missile interplanétaire"))

	// br
	assert.Equal(t, RocketLauncherID, DefenceName2ID("Lançador de Mísseis"))
	assert.Equal(t, LightLaserID, DefenceName2ID("Laser Ligeiro"))
	assert.Equal(t, HeavyLaserID, DefenceName2ID("Laser Pesado"))
	assert.Equal(t, GaussCannonID, DefenceName2ID("Canhão de Gauss"))
	assert.Equal(t, IonCannonID, DefenceName2ID("Canhão de Íons"))
	assert.Equal(t, PlasmaTurretID, DefenceName2ID("Canhão de Plasma"))
	assert.Equal(t, SmallShieldDomeID, DefenceName2ID("Pequeno Escudo Planetário"))
	assert.Equal(t, LargeShieldDomeID, DefenceName2ID("Grande Escudo Planetário"))
	assert.Equal(t, AntiBallisticMissilesID, DefenceName2ID("Míssil de Interceptação"))
	assert.Equal(t, InterplanetaryMissilesID, DefenceName2ID("Míssil Interplanetário"))

	// jp
	assert.Equal(t, RocketLauncherID, DefenceName2ID("ロケットランチャー"))
	assert.Equal(t, LightLaserID, DefenceName2ID("ライトレーザー"))
	assert.Equal(t, HeavyLaserID, DefenceName2ID("ヘビーレーザー"))
	assert.Equal(t, GaussCannonID, DefenceName2ID("ガウスキャノン"))
	assert.Equal(t, IonCannonID, DefenceName2ID("イオンキャノン"))
	assert.Equal(t, PlasmaTurretID, DefenceName2ID("プラズマ砲"))
	assert.Equal(t, SmallShieldDomeID, DefenceName2ID("小型シールドドーム"))
	assert.Equal(t, LargeShieldDomeID, DefenceName2ID("大型シールドドーム"))
	assert.Equal(t, AntiBallisticMissilesID, DefenceName2ID("抗弾道ミサイル"))
	assert.Equal(t, InterplanetaryMissilesID, DefenceName2ID("星間ミサイル"))

	// pl
	assert.Equal(t, RocketLauncherID, DefenceName2ID("Wyrzutnia rakiet"))
	assert.Equal(t, LightLaserID, DefenceName2ID("Lekkie działo laserowe"))
	assert.Equal(t, HeavyLaserID, DefenceName2ID("Ciężkie działo laserowe"))
	assert.Equal(t, GaussCannonID, DefenceName2ID("Działo Gaussa"))
	assert.Equal(t, IonCannonID, DefenceName2ID("Działo jonowe"))
	assert.Equal(t, PlasmaTurretID, DefenceName2ID("Wyrzutnia plazmy"))
	assert.Equal(t, SmallShieldDomeID, DefenceName2ID("Mała Osłona Ochronna"))
	assert.Equal(t, LargeShieldDomeID, DefenceName2ID("Duża Osłona Ochronna"))
	assert.Equal(t, AntiBallisticMissilesID, DefenceName2ID("Przeciwrakieta"))
	assert.Equal(t, InterplanetaryMissilesID, DefenceName2ID("Rakieta międzyplanetarna"))

	// tr
	assert.Equal(t, RocketLauncherID, DefenceName2ID("Roketatar"))
	assert.Equal(t, LightLaserID, DefenceName2ID("Hafif Lazer Topu"))
	assert.Equal(t, HeavyLaserID, DefenceName2ID("Ağır Lazer Topu"))
	assert.Equal(t, GaussCannonID, DefenceName2ID("Gaus Topu"))
	assert.Equal(t, IonCannonID, DefenceName2ID("Iyon Topu"))
	assert.Equal(t, PlasmaTurretID, DefenceName2ID("Plazma Atıcı"))
	assert.Equal(t, SmallShieldDomeID, DefenceName2ID("Küçük Kalkan Kubbesi"))
	assert.Equal(t, LargeShieldDomeID, DefenceName2ID("Büyük Kalkan Kubbesi"))
	assert.Equal(t, AntiBallisticMissilesID, DefenceName2ID("Yakalıyıcı Roketler"))
	assert.Equal(t, InterplanetaryMissilesID, DefenceName2ID("Gezegenlerarasi Roketler"))

	// pt
	assert.Equal(t, RocketLauncherID, DefenceName2ID("Lançador de Mísseis"))
	assert.Equal(t, LightLaserID, DefenceName2ID("Laser Ligeiro"))
	assert.Equal(t, HeavyLaserID, DefenceName2ID("Laser Pesado"))
	assert.Equal(t, GaussCannonID, DefenceName2ID("Canhão de Gauss"))
	assert.Equal(t, IonCannonID, DefenceName2ID("Canhão de Iões"))
	assert.Equal(t, PlasmaTurretID, DefenceName2ID("Canhão de Plasma"))
	assert.Equal(t, SmallShieldDomeID, DefenceName2ID("Pequeno Escudo Planetário"))
	assert.Equal(t, LargeShieldDomeID, DefenceName2ID("Grande Escudo Planetário"))
	assert.Equal(t, AntiBallisticMissilesID, DefenceName2ID("Míssil de Intercepção"))
	assert.Equal(t, InterplanetaryMissilesID, DefenceName2ID("Míssil Interplanetário"))

	// nl
	assert.Equal(t, RocketLauncherID, DefenceName2ID("Raketlanceerder"))
	assert.Equal(t, LightLaserID, DefenceName2ID("Kleine laser"))
	assert.Equal(t, HeavyLaserID, DefenceName2ID("Grote laser"))
	assert.Equal(t, GaussCannonID, DefenceName2ID("Gausskanon"))
	assert.Equal(t, IonCannonID, DefenceName2ID("Ionkanon"))
	assert.Equal(t, PlasmaTurretID, DefenceName2ID("Plasmakanon"))
	assert.Equal(t, SmallShieldDomeID, DefenceName2ID("Kleine planetaire schildkoepel"))
	assert.Equal(t, LargeShieldDomeID, DefenceName2ID("Grote planetaire schildkoepel"))
	assert.Equal(t, AntiBallisticMissilesID, DefenceName2ID("Anti-ballistische raketten"))
	assert.Equal(t, InterplanetaryMissilesID, DefenceName2ID("Interplanetaire raketten"))

	// ru
	assert.Equal(t, RocketLauncherID, DefenceName2ID("Ракетная установка"))
	assert.Equal(t, LightLaserID, DefenceName2ID("Лёгкий лазер"))
	assert.Equal(t, HeavyLaserID, DefenceName2ID("Тяжёлый лазер"))
	assert.Equal(t, GaussCannonID, DefenceName2ID("Пушка Гаусса"))
	assert.Equal(t, IonCannonID, DefenceName2ID("Ионное орудие"))
	assert.Equal(t, PlasmaTurretID, DefenceName2ID("Плазменное орудие"))
	assert.Equal(t, SmallShieldDomeID, DefenceName2ID("Малый щитовой купол"))
	assert.Equal(t, LargeShieldDomeID, DefenceName2ID("Большой щитовой купол"))
	assert.Equal(t, AntiBallisticMissilesID, DefenceName2ID("Ракета-перехватчик"))
	assert.Equal(t, InterplanetaryMissilesID, DefenceName2ID("Межпланетная ракета"))

	// SHIPS
	assert.Equal(t, ID(0), ShipName2ID("Not valid"))
	assert.Equal(t, LightFighterID, ShipName2ID("Light Fighter"))
	assert.Equal(t, LightFighterID, ShipName2ID("Chasseur léger"))
	assert.Equal(t, LightFighterID, ShipName2ID("Leichter Jäger"))
	assert.Equal(t, LightFighterID, ShipName2ID("Caça Ligeiro"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Caça Pesado"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Großer Transporter"))
	assert.Equal(t, DestroyerID, ShipName2ID("Zerstörer"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Nave pequeña de carga"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Satélite solar"))
	assert.Equal(t, ID(0), ShipName2ID("人中位"))

	// fi
	assert.Equal(t, LightFighterID, ShipName2ID("Kevyt Hävittäjä"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Raskas Hävittäjä"))
	assert.Equal(t, CruiserID, ShipName2ID("Risteilijä"))
	assert.Equal(t, BattleshipID, ShipName2ID("Taistelualus"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Taisteluristeilijä"))
	assert.Equal(t, BomberID, ShipName2ID("Pommittaja"))
	assert.Equal(t, DestroyerID, ShipName2ID("Tuhoaja"))
	assert.Equal(t, DeathstarID, ShipName2ID("Kuolemantähti"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Pieni rahtialus"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Suuri rahtialus"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Siirtokunta-alus"))
	assert.Equal(t, RecyclerID, ShipName2ID("Kierrättäjä"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Vakoiluluotain"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Aurinkosatelliitti"))
	assert.Equal(t, CrawlerID, ShipName2ID("Crawler"))
	assert.Equal(t, ReaperID, ShipName2ID("Reaper"))
	assert.Equal(t, PathfinderID, ShipName2ID("Pathfinder"))

	// hu
	assert.Equal(t, LightFighterID, ShipName2ID("Könnyű Harcos"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Nehéz Harcos"))
	assert.Equal(t, CruiserID, ShipName2ID("Cirkáló"))
	assert.Equal(t, BattleshipID, ShipName2ID("Csatahajó"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Csatacirkáló"))
	assert.Equal(t, BomberID, ShipName2ID("Bombázó"))
	assert.Equal(t, DestroyerID, ShipName2ID("Romboló"))
	assert.Equal(t, DeathstarID, ShipName2ID("Halálcsillag"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Kis szállító"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Nagy Szállító"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Kolóniahajó"))
	assert.Equal(t, RecyclerID, ShipName2ID("Szemetesek"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Kémszonda"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Napműhold"))
	assert.Equal(t, CrawlerID, ShipName2ID("Crawler"))
	assert.Equal(t, ReaperID, ShipName2ID("Kaszás"))
	assert.Equal(t, PathfinderID, ShipName2ID("Felderítő"))

	// ro
	assert.Equal(t, LightFighterID, ShipName2ID("Vanator usor"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Vanator greu"))
	assert.Equal(t, CruiserID, ShipName2ID("Crucisator"))
	assert.Equal(t, BattleshipID, ShipName2ID("Nava de razboi"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Interceptor"))
	assert.Equal(t, BomberID, ShipName2ID("Bombardier"))
	assert.Equal(t, DestroyerID, ShipName2ID("Distrugator"))
	assert.Equal(t, DeathstarID, ShipName2ID("RIP"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Transportor mic"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Transportor mare"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Nava de Colonizare"))
	assert.Equal(t, RecyclerID, ShipName2ID("Reciclator"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Proba de spionaj"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Satelit solar"))
	assert.Equal(t, CrawlerID, ShipName2ID("Crawler"))
	assert.Equal(t, ReaperID, ShipName2ID("Reaper"))
	assert.Equal(t, PathfinderID, ShipName2ID("Pathfinder"))

	// cz
	assert.Equal(t, LightFighterID, ShipName2ID("Lehký stíhač"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Těžký stíhač"))
	assert.Equal(t, CruiserID, ShipName2ID("Křižník"))
	assert.Equal(t, BattleshipID, ShipName2ID("Bitevní loď"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Bitevní křižník"))
	assert.Equal(t, BomberID, ShipName2ID("Bombardér"))
	assert.Equal(t, DestroyerID, ShipName2ID("Ničitel"))
	assert.Equal(t, DeathstarID, ShipName2ID("Hvězda smrti"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Malý transportér"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Velký transportér"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Kolonizační loď"))
	assert.Equal(t, RecyclerID, ShipName2ID("Recyklátor"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Špionážní sonda"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Solární satelit"))
	assert.Equal(t, CrawlerID, ShipName2ID("Crawler"))
	assert.Equal(t, ReaperID, ShipName2ID("Rozparovač"))
	assert.Equal(t, PathfinderID, ShipName2ID("Průzkumník"))

	// mx
	assert.Equal(t, LightFighterID, ShipName2ID("Cazador ligero"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Cazador pesado"))
	assert.Equal(t, CruiserID, ShipName2ID("Crucero"))
	assert.Equal(t, BattleshipID, ShipName2ID("Nave de batalla"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Acorazado"))
	assert.Equal(t, BomberID, ShipName2ID("Bombardero"))
	assert.Equal(t, DestroyerID, ShipName2ID("Destructor"))
	assert.Equal(t, DeathstarID, ShipName2ID("Estrella de la muerte"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Nave pequeña de carga"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Nave grande de carga"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Nave de la colonia"))
	assert.Equal(t, RecyclerID, ShipName2ID("Reciclador"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Sonda de espionaje"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Satélite solar"))
	assert.Equal(t, CrawlerID, ShipName2ID("Taladrador"))
	assert.Equal(t, ReaperID, ShipName2ID("Segador"))
	assert.Equal(t, PathfinderID, ShipName2ID("Explorador"))

	// hr
	assert.Equal(t, LightFighterID, ShipName2ID("Mali lovac"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Veliki lovac"))
	assert.Equal(t, CruiserID, ShipName2ID("Krstarica"))
	assert.Equal(t, BattleshipID, ShipName2ID("Borbeni brod"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Oklopna krstarica"))
	assert.Equal(t, BomberID, ShipName2ID("Bombarder"))
	assert.Equal(t, DestroyerID, ShipName2ID("Razarač"))
	assert.Equal(t, DeathstarID, ShipName2ID("Zvijezda smrti"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Mali transporter"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Veliki transporter"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Kolonijalni brod"))
	assert.Equal(t, RecyclerID, ShipName2ID("Recikler"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Sonde za špijunažu"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Solarni satelit"))
	assert.Equal(t, CrawlerID, ShipName2ID("Puzavac"))
	assert.Equal(t, ReaperID, ShipName2ID("Žetelac"))
	assert.Equal(t, PathfinderID, ShipName2ID("Krčilac"))

	// ba
	assert.Equal(t, LightFighterID, ShipName2ID("Mali lovac"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Veliki lovac"))
	assert.Equal(t, CruiserID, ShipName2ID("Krstarice"))
	assert.Equal(t, BattleshipID, ShipName2ID("Borbeni brodovi"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Oklopna krstarica"))
	assert.Equal(t, BomberID, ShipName2ID("Bombarder"))
	assert.Equal(t, DestroyerID, ShipName2ID("Razaraci"))
	assert.Equal(t, DeathstarID, ShipName2ID("Zvijezda smrti"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Mali transporter"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Veliki transporter"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Kolonijalni brod"))
	assert.Equal(t, RecyclerID, ShipName2ID("Recikler"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Sonde za spijunazu"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Solarni satelit"))
	assert.Equal(t, CrawlerID, ShipName2ID("Puzavac"))
	assert.Equal(t, ReaperID, ShipName2ID("Žetelac"))
	assert.Equal(t, PathfinderID, ShipName2ID("Krčilac"))

	// no
	assert.Equal(t, LightFighterID, ShipName2ID("Lett Jeger"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Tung Jeger"))
	assert.Equal(t, CruiserID, ShipName2ID("Krysser"))
	assert.Equal(t, BattleshipID, ShipName2ID("Slagskip"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Slagkrysser"))
	assert.Equal(t, BomberID, ShipName2ID("Bomber"))
	assert.Equal(t, DestroyerID, ShipName2ID("Destroyer"))
	assert.Equal(t, DeathstarID, ShipName2ID("Døds stjerne"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Lite Lasteskip"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Stort Lasteskip"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Koloni Skip"))
	assert.Equal(t, RecyclerID, ShipName2ID("Resirkulerer"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Spionasjesonde"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Solar Satelitt"))
	assert.Equal(t, CrawlerID, ShipName2ID("Crawler"))
	assert.Equal(t, ReaperID, ShipName2ID("Reaper"))
	assert.Equal(t, PathfinderID, ShipName2ID("Pathfinder"))

	// it
	assert.Equal(t, LightFighterID, ShipName2ID("Caccia Leggero"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Caccia Pesante"))
	assert.Equal(t, CruiserID, ShipName2ID("Incrociatore"))
	assert.Equal(t, BattleshipID, ShipName2ID("Nave da battaglia"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Incrociatore da Battaglia"))
	assert.Equal(t, BomberID, ShipName2ID("Bombardiere"))
	assert.Equal(t, DestroyerID, ShipName2ID("Corazzata"))
	assert.Equal(t, DeathstarID, ShipName2ID("Morte Nera"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Cargo leggero"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Cargo Pesante"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Colonizzatrice"))
	assert.Equal(t, RecyclerID, ShipName2ID("Riciclatrici"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Sonda spia"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Satellite Solare"))
	assert.Equal(t, CrawlerID, ShipName2ID("Crawler"))
	assert.Equal(t, ReaperID, ShipName2ID("Reaper"))
	assert.Equal(t, PathfinderID, ShipName2ID("Pathfinder"))

	// pl
	assert.Equal(t, LightFighterID, ShipName2ID("Lekki myśliwiec"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Ciężki myśliwiec"))
	assert.Equal(t, CruiserID, ShipName2ID("Krążownik"))
	assert.Equal(t, BattleshipID, ShipName2ID("Okręt wojenny"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Pancernik"))
	assert.Equal(t, BomberID, ShipName2ID("Bombowiec"))
	assert.Equal(t, DestroyerID, ShipName2ID("Niszczyciel"))
	assert.Equal(t, DeathstarID, ShipName2ID("Gwiazda Śmierci"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Mały transporter"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Duży transporter"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Statek kolonizacyjny"))
	assert.Equal(t, RecyclerID, ShipName2ID("Recykler"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Sonda szpiegowska"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Satelita słoneczny"))
	assert.Equal(t, CrawlerID, ShipName2ID("Pełzacz"))
	assert.Equal(t, ReaperID, ShipName2ID("Rozpruwacz"))
	assert.Equal(t, PathfinderID, ShipName2ID("Pionier"))

	// tr
	assert.Equal(t, LightFighterID, ShipName2ID("Hafif Avcı"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Ağır Avcı"))
	assert.Equal(t, CruiserID, ShipName2ID("Kruvazör"))
	assert.Equal(t, BattleshipID, ShipName2ID("Komuta Gemisi"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Firkateyn"))
	assert.Equal(t, BomberID, ShipName2ID("Bombardıman Gemisi"))
	assert.Equal(t, DestroyerID, ShipName2ID("Muhrip"))
	assert.Equal(t, DeathstarID, ShipName2ID("Ölüm Yildizi"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Küçük Nakliye Gemisi"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Büyük Nakliye Gemisi"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Koloni Gemisi"))
	assert.Equal(t, RecyclerID, ShipName2ID("Geri Dönüsümcü"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Casus Sondasi"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Solar Uydu"))
	assert.Equal(t, CrawlerID, ShipName2ID("Paletli"))
	assert.Equal(t, ReaperID, ShipName2ID("Azrail"))
	assert.Equal(t, PathfinderID, ShipName2ID("Rehber"))

	// ar
	assert.Equal(t, LightFighterID, ShipName2ID("Cazador ligero"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Cazador pesado"))
	assert.Equal(t, CruiserID, ShipName2ID("Crucero"))
	assert.Equal(t, BattleshipID, ShipName2ID("Nave de batalla"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Acorazado"))
	assert.Equal(t, BomberID, ShipName2ID("Bombardero"))
	assert.Equal(t, DestroyerID, ShipName2ID("Destructor"))
	assert.Equal(t, DeathstarID, ShipName2ID("Estrella de la muerte"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Nave pequeña de carga"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Nave grande de carga"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Nave colonizadora"))
	assert.Equal(t, RecyclerID, ShipName2ID("Reciclador"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Sonda de espionaje"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Satélite solar"))
	assert.Equal(t, CrawlerID, ShipName2ID("Taladrador"))
	assert.Equal(t, ReaperID, ShipName2ID("Segador"))
	assert.Equal(t, PathfinderID, ShipName2ID("Explorador"))

	// pt
	assert.Equal(t, LightFighterID, ShipName2ID("Caça Ligeiro"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Caça Pesado"))
	assert.Equal(t, CruiserID, ShipName2ID("Cruzador"))
	assert.Equal(t, BattleshipID, ShipName2ID("Nave de Batalha"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Interceptor"))
	assert.Equal(t, BomberID, ShipName2ID("Bombardeiro"))
	assert.Equal(t, DestroyerID, ShipName2ID("Destruidor"))
	assert.Equal(t, DeathstarID, ShipName2ID("Estrela da Morte"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Cargueiro Pequeno"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Cargueiro Grande"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Nave de Colonização"))
	assert.Equal(t, RecyclerID, ShipName2ID("Reciclador"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Sonda de Espionagem"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Satélite Solar"))
	assert.Equal(t, CrawlerID, ShipName2ID("Rastejador"))
	assert.Equal(t, ReaperID, ShipName2ID("Ceifeira"))
	assert.Equal(t, PathfinderID, ShipName2ID("Exploradora"))

	// nl
	assert.Equal(t, LightFighterID, ShipName2ID("Licht gevechtsschip"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Zwaar gevechtsschip"))
	assert.Equal(t, CruiserID, ShipName2ID("Kruiser"))
	assert.Equal(t, BattleshipID, ShipName2ID("Slagschip"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Interceptor"))
	assert.Equal(t, BomberID, ShipName2ID("Bommenwerper"))
	assert.Equal(t, DestroyerID, ShipName2ID("Vernietiger"))
	assert.Equal(t, DeathstarID, ShipName2ID("Ster des Doods"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Klein vrachtschip"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Groot vrachtschip"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Kolonisatieschip"))
	assert.Equal(t, RecyclerID, ShipName2ID("Recycler"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Spionagesonde"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Zonne-energiesatelliet"))
	assert.Equal(t, CrawlerID, ShipName2ID("Processer"))
	assert.Equal(t, ReaperID, ShipName2ID("Ruimer"))
	assert.Equal(t, PathfinderID, ShipName2ID("Navigator"))

	// dk
	assert.Equal(t, LightFighterID, ShipName2ID("Lille Jæger"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Stor Jæger"))
	assert.Equal(t, CruiserID, ShipName2ID("Krydser"))
	assert.Equal(t, BattleshipID, ShipName2ID("Slagskib"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Interceptor"))
	assert.Equal(t, BomberID, ShipName2ID("Bomber"))
	assert.Equal(t, DestroyerID, ShipName2ID("Destroyer"))
	assert.Equal(t, DeathstarID, ShipName2ID("Dødsstjerne"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Lille Transporter"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Stor Transporter"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Koloniskib"))
	assert.Equal(t, RecyclerID, ShipName2ID("Recycler"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Spionagesonde"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Solarsatellit"))
	assert.Equal(t, CrawlerID, ShipName2ID("Kravler"))
	assert.Equal(t, ReaperID, ShipName2ID("Reaper"))
	assert.Equal(t, PathfinderID, ShipName2ID("Stifinder"))

	// ru
	assert.Equal(t, LightFighterID, ShipName2ID("Лёгкий истребитель"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Тяжёлый истребитель"))
	assert.Equal(t, CruiserID, ShipName2ID("Крейсер"))
	assert.Equal(t, BattleshipID, ShipName2ID("Линкор"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Линейный крейсер"))
	assert.Equal(t, BomberID, ShipName2ID("Бомбардировщик"))
	assert.Equal(t, DestroyerID, ShipName2ID("Уничтожитель"))
	assert.Equal(t, DeathstarID, ShipName2ID("Звезда смерти"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Малый транспорт"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Большой транспорт"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Колонизатор"))
	assert.Equal(t, RecyclerID, ShipName2ID("Переработчик"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Шпионский зонд"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Солнечный спутник"))
	assert.Equal(t, CrawlerID, ShipName2ID("Гусеничник"))
	assert.Equal(t, ReaperID, ShipName2ID("Жнец"))
	assert.Equal(t, PathfinderID, ShipName2ID("Первопроходец"))

	// gr
	assert.Equal(t, LightFighterID, ShipName2ID("Ελαφρύ Μαχητικό"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Βαρύ Μαχητικό"))
	assert.Equal(t, CruiserID, ShipName2ID("Καταδιωκτικό"))
	assert.Equal(t, BattleshipID, ShipName2ID("Καταδρομικό"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Θωρηκτό Αναχαίτισης"))
	assert.Equal(t, BomberID, ShipName2ID("Βομβαρδιστικό"))
	assert.Equal(t, DestroyerID, ShipName2ID("Destroyer"))
	assert.Equal(t, DeathstarID, ShipName2ID("Deathstar"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Μικρό Μεταγωγικό"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Μεγάλο Μεταγωγικό"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Σκάφος Αποικιοποίησης"))
	assert.Equal(t, RecyclerID, ShipName2ID("Ανακυκλωτής"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Κατασκοπευτικό Στέλεχος"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Ηλιακοί Συλλέκτες"))
	assert.Equal(t, CrawlerID, ShipName2ID("Crawler"))
	assert.Equal(t, ReaperID, ShipName2ID("Reaper"))
	assert.Equal(t, PathfinderID, ShipName2ID("Pathfinder"))

	// jp
	assert.Equal(t, LightFighterID, ShipName2ID("軽戦闘機"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("重戦闘機"))
	assert.Equal(t, CruiserID, ShipName2ID("巡洋艦"))
	assert.Equal(t, BattleshipID, ShipName2ID("バトルシップ"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("大型戦艦"))
	assert.Equal(t, BomberID, ShipName2ID("爆撃機"))
	assert.Equal(t, DestroyerID, ShipName2ID("デストロイヤー"))
	assert.Equal(t, DeathstarID, ShipName2ID("デススター"))
	assert.Equal(t, SmallCargoID, ShipName2ID("小型輸送機"))
	assert.Equal(t, LargeCargoID, ShipName2ID("大型輸送機"))
	assert.Equal(t, ColonyShipID, ShipName2ID("コロニーシップ"))
	assert.Equal(t, RecyclerID, ShipName2ID("残骸回収船"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("偵察機"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("ソーラーサテライト"))
	assert.Equal(t, CrawlerID, ShipName2ID("クローラー"))
	assert.Equal(t, ReaperID, ShipName2ID("リーパー"))
	assert.Equal(t, PathfinderID, ShipName2ID("パスファインダー"))

	// sk
	assert.Equal(t, LightFighterID, ShipName2ID("Ľahký stíhač"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Ťažký stíhač"))
	assert.Equal(t, CruiserID, ShipName2ID("Krížnik"))
	assert.Equal(t, BattleshipID, ShipName2ID("Bojová loď"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Bojový krížnik"))
	assert.Equal(t, BomberID, ShipName2ID("Bombardér"))
	assert.Equal(t, DestroyerID, ShipName2ID("Devastátor"))
	assert.Equal(t, DeathstarID, ShipName2ID("Hviezda smrti"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Malý transportér"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Veľký transportér"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Kolonizačná loď"))
	assert.Equal(t, RecyclerID, ShipName2ID("Recyklátor"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Špionážna sonda"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Solárny satelit"))
	assert.Equal(t, CrawlerID, ShipName2ID("Vrták"))
	assert.Equal(t, ReaperID, ShipName2ID("Kosa"))
	assert.Equal(t, PathfinderID, ShipName2ID("Prieskumník"))

	// si
	assert.Equal(t, LightFighterID, ShipName2ID("Lahek lovec"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Težki lovec"))
	assert.Equal(t, CruiserID, ShipName2ID("Križarka"))
	assert.Equal(t, BattleshipID, ShipName2ID("Bojna ladja"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Bojna križarka"))
	assert.Equal(t, BomberID, ShipName2ID("Bombnik"))
	assert.Equal(t, DestroyerID, ShipName2ID("Uničevalec"))
	assert.Equal(t, DeathstarID, ShipName2ID("Zvezda smrti"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Majhna tovorna ladja"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Velika tovorna ladja"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Kolonizacijska ladja"))
	assert.Equal(t, RecyclerID, ShipName2ID("Recikler"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Vohunska sonda"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Sončni satelit"))
	assert.Equal(t, CrawlerID, ShipName2ID("Plazilec"))
	assert.Equal(t, ReaperID, ShipName2ID("Kombajn"))
	assert.Equal(t, PathfinderID, ShipName2ID("Iskalec sledi"))

	// fr
	assert.Equal(t, LightFighterID, ShipName2ID("Chasseur léger"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Chasseur lourd"))
	assert.Equal(t, CruiserID, ShipName2ID("Croiseur"))
	assert.Equal(t, BattleshipID, ShipName2ID("Vaisseau de bataille"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Traqueur"))
	assert.Equal(t, BomberID, ShipName2ID("Bombardier"))
	assert.Equal(t, DestroyerID, ShipName2ID("Destructeur"))
	assert.Equal(t, DeathstarID, ShipName2ID("Étoile de la mort"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Petit transporteur"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Grand transporteur"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Vaisseau de colonisation"))
	assert.Equal(t, RecyclerID, ShipName2ID("Recycleur"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Sonde d`espionnage"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Satellite solaire"))
	assert.Equal(t, CrawlerID, ShipName2ID("Foreuse"))
	assert.Equal(t, ReaperID, ShipName2ID("Faucheur"))
	assert.Equal(t, PathfinderID, ShipName2ID("Éclaireur"))

	// tw
	assert.Equal(t, LightFighterID, ShipName2ID("輕型戰鬥機"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("重型戰鬥機"))
	assert.Equal(t, CruiserID, ShipName2ID("巡洋艦"))
	assert.Equal(t, BattleshipID, ShipName2ID("戰列艦"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("戰鬥巡洋艦"))
	assert.Equal(t, BomberID, ShipName2ID("導彈艦"))
	assert.Equal(t, DestroyerID, ShipName2ID("毀滅者"))
	assert.Equal(t, DeathstarID, ShipName2ID("死星"))
	assert.Equal(t, SmallCargoID, ShipName2ID("小型運輸艦"))
	assert.Equal(t, LargeCargoID, ShipName2ID("大型運輸艦"))
	assert.Equal(t, ColonyShipID, ShipName2ID("殖民船"))
	assert.Equal(t, RecyclerID, ShipName2ID("回收船"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("間諜衛星"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("太陽能衛星"))
	assert.Equal(t, CrawlerID, ShipName2ID("履帶車"))
	assert.Equal(t, ReaperID, ShipName2ID("惡魔飛船"))
	assert.Equal(t, PathfinderID, ShipName2ID("探路者"))
}
