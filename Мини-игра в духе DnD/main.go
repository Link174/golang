package main

/*Запустите код через терминал*/

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	playerTakeInit bool   = false
	playerLive     bool   = true
	playerRace     string = "Человек"
	playerClass    string = "Воин"
	playerLev      int    = 1
	playerExp      int    = 0
	playerDef      int    = 14
	playerHP       int    = 20
	playerStr      int    = 4
	playerDex      int    = 2
	playerCon      int    = 3
	playerInt      int    = 1
	playerWis      int    = -2
	playerChar     int    = 2
	playerCoins    int    = 100
)

var (
	enemyTakeInit bool   = false
	enemyLive     bool   = true
	enemyRace     string = ""
	enemyClass    string = ""
	enemyDef      int    = 0
	enemyHP       int    = 0
	enemyStr      int    = 0
	enemyDex      int    = 0
	enemyCon      int    = 0
	enemyInt      int    = 0
	enemyWis      int    = 0
	enemyChar     int    = 0
)

var diceResult int = 0

func dice20() int {
	diceResult = rand.Intn(20)
	diceResult += 1
	return diceResult
}

func dice10() int {
	diceResult = rand.Intn(10)
	diceResult += 1
	return diceResult
}

func dice6() int {
	diceResult = rand.Intn(6)
	diceResult += 1
	return diceResult
}

func intro() {
	fmt.Println("Привет! Эта мини-игра вдохновлена системой DnD. Ничего делать не нужно, всё происходит само ;). Для удобства разверните терминал побольше. Желаю критической удачи!")

	time.Sleep(8000 * time.Millisecond)

	fmt.Println()

	fmt.Println("Вы в городе Rosewater. Это мелкий преступный город на окраине региона Lowland. Когда-то этот город процветал, ведь здесь проходил главный торговый путь через горы.")

	time.Sleep(8000 * time.Millisecond)

	fmt.Println()

	fmt.Println("Город стал известен своей торговлей так называемым розовым вином.")

	time.Sleep(6000 * time.Millisecond)

	fmt.Println()

	fmt.Println("Но однажды в горе поселился дракон... От этого торговый путь стал очень опасен. Спустя годы от былого процветания остались лишь шайки преступников и контрабандистов.")

	time.Sleep(8000 * time.Millisecond)

	fmt.Println()

	fmt.Println("Ваши цели в этом городе предельно ясны. Вас отправили разделаться с главарём контрабандистов. Но это ещё та задача...")

	time.Sleep(6500 * time.Millisecond)

	fmt.Println()

	fmt.Println("Мало того, что весь город до отвала забит бедняками да преступниками, в нём ещё и с недавних пор стали создавать огромных стальных големов.")

	time.Sleep(8000 * time.Millisecond)

	fmt.Println()

	fmt.Println("Так контрабандисты наращивают свою мощь и в то же время зарабатывают, тайно торгуя с соседним королевством...")

	time.Sleep(6500 * time.Millisecond)

	fmt.Println()

	fmt.Println("Именно поэтому вы здесь - спасите свои родные земли!")

	time.Sleep(2000 * time.Millisecond)

	fmt.Println("________________________________________________________________________________")

	time.Sleep(4000 * time.Millisecond)

	fmt.Println()

	fmt.Println("Ознакомьтесь с вашими характеристиками:")

	time.Sleep(2500 * time.Millisecond)

	fmt.Println()

	fmt.Println("Раса:", playerRace)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Класс:", playerClass)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Уровень:", playerLev)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Очки опыта:", playerExp)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Очки защиты:", playerDef)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Очки здоровья:", playerHP)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Сила:", playerStr)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Ловкость:", playerDex)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Выносливость:", playerCon)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Интеллект:", playerInt)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Мудрость:", playerWis)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Харизма:", playerChar)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Ваши деньги:", playerCoins)

	time.Sleep(500 * time.Millisecond)

	fmt.Println("________________________________________________________________________________")

	time.Sleep(12000 * time.Millisecond)

	fmt.Println()

	fmt.Println("Итак, вы идёте по узкой улочке города. Вы замотаны в ту же одежду, что и местные преступники...")

	time.Sleep(6000 * time.Millisecond)
}

func spawnEnemy() {
	time.Sleep(4000 * time.Millisecond)

	fmt.Println()

	var enemy string

	spawnEnemy := rand.Intn(4)

	if spawnEnemy == 0 {
		enemy = "Из-за угла появился вооружённый мужчина! Его модификатор защиты: 13. У него 13 очков жизни"
		enemyLive = true
		enemyClass = "Воин"
		enemyRace = "Человек"
		enemyDef = 13
		enemyHP = 13
		enemyStr = 2
		enemyDex = 2
		enemyCon = 1
		enemyInt = 1
		enemyWis = 0
		enemyChar = 1
	} else if spawnEnemy == 1 {
		enemy = "Что-то зашевелилось в груде металла... Это железный голем! Его модификатор защиты: 16. У него 16 очков жизни"
		enemyLive = true
		enemyClass = "Воин"
		enemyRace = "-"
		enemyDef = 16
		enemyHP = 16
		enemyStr = 4
		enemyDex = 0
		enemyCon = 3
		enemyInt = 0
		enemyWis = -1
		enemyChar = -2
	} else if spawnEnemy == 2 {
		enemy = "Воздух впереди сгущается... Перед тобой появляется убийца! Его модификатор защиты: 15. У него 14 очков жизни"
		enemyLive = true
		enemyClass = "Следопыт"
		enemyRace = "Тёмный дроу"
		enemyDef = 15
		enemyHP = 14
		enemyStr = 3
		enemyDex = 2
		enemyCon = 1
		enemyInt = 4
		enemyWis = 5
		enemyChar = 2
	} else if spawnEnemy == 3 {
		enemy = "Ты замечаешь какое-то движение в тени обветшалого здания... На тебя побежал безумный мужчина! Его модификатор защиты: 13. У него 11 очков жизни"
		enemyLive = true
		enemyClass = "Воин"
		enemyRace = "Полуэльф"
		enemyDef = 13
		enemyHP = 11
		enemyStr = 2
		enemyDex = 2
		enemyCon = 1
		enemyInt = 1
		enemyWis = 0
		enemyChar = 1
	}

	fmt.Println(enemy)

	time.Sleep(6000 * time.Millisecond)
}

func whoInit() {
	fmt.Println()

	playerDice := dice20()

	enemyDice := dice20()

	fmt.Println("Определим инициативу. Бросаю кубики...")

	fmt.Println()

	time.Sleep(4000 * time.Millisecond)

	if playerDice > enemyDice {
		fmt.Println("Инициатива за тобой!", playerDice, "против", enemyDice, ".")
		playerTakeInit = true
	} else if playerDice < enemyDice {
		fmt.Println("Инициатива за противником!", enemyDice, "против", playerDice, ".")
		enemyTakeInit = true
	} else {
		fmt.Println("Никто не принял инициативу... Ход пропущен.")
	}
}

func playerAttack() {
	time.Sleep(3000 * time.Millisecond)

	fmt.Println()

	fmt.Println("Ты бросаешь кубик на атаку...")

	time.Sleep(3000 * time.Millisecond)

	fmt.Println()

	a := dice20()

	if a == 20 {
		fmt.Println("Критическая удача!")
	} else if a == 1 {
		fmt.Println("Критический провал!")
	} else {
		fmt.Println("Выпало:", a)
	}

	time.Sleep(2000 * time.Millisecond)

	fmt.Println()

	fmt.Println("Плюс твой модификатор силы...")

	time.Sleep(3000 * time.Millisecond)

	fmt.Println()

	result := a + playerStr

	fmt.Println("Результат:", result)

	time.Sleep(2000 * time.Millisecond)

	fmt.Println()

	if result >= enemyDef {
		fmt.Println("Попадание!")

		time.Sleep(2000 * time.Millisecond)

		fmt.Println()

		fmt.Println("Определим урон...")

		time.Sleep(2000 * time.Millisecond)

		fmt.Println()

		b := dice10()

		if b == 20 {
			fmt.Println("Критическая удача!")
		} else if b == 1 {
			fmt.Println("Критический провал!")
		} else {
			fmt.Println("Выпало:", b)
		}

		time.Sleep(2000 * time.Millisecond)

		fmt.Println()

		fmt.Println("Плюс твой модификатор силы...")

		time.Sleep(3000 * time.Millisecond)

		fmt.Println()

		damage := b + playerStr

		fmt.Println("Ты наносишь", damage, "урона.")

		enemyHP -= damage
	} else {
		fmt.Println("Промах!")

	}
}

func enemyAttack() {
	time.Sleep(3000 * time.Millisecond)

	fmt.Println()

	fmt.Println("Противник бросает кубик на атаку...")

	time.Sleep(3000 * time.Millisecond)

	fmt.Println()

	a := dice20()

	if a == 20 {
		fmt.Println("Критическая удача!")
	} else if a == 1 {
		fmt.Println("Критический провал!")
	} else {
		fmt.Println("Выпало:", a)
	}

	time.Sleep(2000 * time.Millisecond)

	fmt.Println()

	fmt.Println("Плюс его модификатор силы...")

	time.Sleep(3000 * time.Millisecond)

	fmt.Println()

	result := a + enemyStr

	fmt.Println("Результат:", result)

	time.Sleep(2000 * time.Millisecond)

	fmt.Println()

	if result >= playerDef {
		fmt.Println("Попадание!")
		time.Sleep(2000 * time.Millisecond)

		fmt.Println()

		fmt.Println("Кидаю на урон...")

		time.Sleep(2000 * time.Millisecond)

		fmt.Println()

		b := dice10()

		if b == 20 {
			fmt.Println("Критическая удача!")
		} else if b == 1 {
			fmt.Println("Критический провал!")
		} else {
			fmt.Println("Выпало:", b)
		}

		time.Sleep(2000 * time.Millisecond)

		fmt.Println()

		fmt.Println("Плюс его модификатор силы...")

		time.Sleep(3000 * time.Millisecond)

		fmt.Println()

		damage := b + enemyStr

		fmt.Println("Тебе наносят", damage, "урона.")

		playerHP -= damage
	} else {
		fmt.Println("Промах!")

	}
}

func battle() {
	for {
		time.Sleep(2000 * time.Millisecond)

		whoInit()

		if playerTakeInit {
			playerAttack()
			playerTakeInit = false
		} else if enemyTakeInit {
			enemyAttack()
			enemyTakeInit = false
		} else if playerTakeInit && enemyTakeInit == false {
			battle()
		}

		if enemyHP < 1 {
			fmt.Println()
			fmt.Println("Враг повержен! Поздравляю!")
			break
		} else if playerHP < 1 {
			fmt.Println("You dead")
			time.Sleep(4000 * time.Millisecond)
			fmt.Println("Не переживайте, о вас будут слогать легенды...")

			time.Sleep(2000 * time.Millisecond)

			fmt.Println()

			fmt.Println("________________________________________________________________________________")
			break
		}
	}
}

func goNext() {
	time.Sleep(3000 * time.Millisecond)

	fmt.Println()

	fmt.Println("Это была хорошая битва, но тебе нужно передохнуть.")

	time.Sleep(4000 * time.Millisecond)

	fmt.Println()

	fmt.Println("Определим сколько очков здоровья ты восстановишь...")

	time.Sleep(4000 * time.Millisecond)

	fmt.Println()

	playerHilling := dice6()

	fmt.Println("Ты восстановил", playerHilling, "очков здоровья.")

	playerHP += playerHilling

	time.Sleep(4000 * time.Millisecond)

	fmt.Println()

	fmt.Println("Набравшись сил, ты направляешься дальше...")

	time.Sleep(2000 * time.Millisecond)

	fmt.Println()

	fmt.Println("________________________________________________________________________________")
}

func main() {

	spawnEnemy()

	battle()

	goNext()

	spawnEnemy()

	battle()

	endGame()
}

func endGame() {
	time.Sleep(4000 * time.Millisecond)

	fmt.Println()

	fmt.Println("Игра завершена... Надеюсь вам понравилось! При желании запустите игру ещё раз.")
}
