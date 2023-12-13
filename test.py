## random python file code 

import random
import statistics

def generate_random_list(size):
    return [random.randint(1, 100) for _ in range(size)]

def calculate_mean_and_std_dev(data):
    mean_value = statistics.mean(data)
    std_dev_value = statistics.stdev(data)
    return mean_value, std_dev_value

def main():
    list_size = 10
    random_numbers = generate_random_list(list_size)

    print(f"Random Numbers: {random_numbers}")

    mean, std_dev = calculate_mean_and_std_dev(random_numbers)

    print(f"Mean: {mean}")
    print(f"Standard Deviation: {std_dev}")

if __name__ == "__main__":
    main()
