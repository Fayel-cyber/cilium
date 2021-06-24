package vpc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Ipv6Address is a nested struct in vpc response
type Ipv6Address struct {
	Ipv6AddressId          string                `json:"Ipv6AddressId" xml:"Ipv6AddressId"`
	Ipv6AddressName        string                `json:"Ipv6AddressName" xml:"Ipv6AddressName"`
	VSwitchId              string                `json:"VSwitchId" xml:"VSwitchId"`
	VpcId                  string                `json:"VpcId" xml:"VpcId"`
	Ipv6GatewayId          string                `json:"Ipv6GatewayId" xml:"Ipv6GatewayId"`
	Ipv6Address            string                `json:"Ipv6Address" xml:"Ipv6Address"`
	AssociatedInstanceId   string                `json:"AssociatedInstanceId" xml:"AssociatedInstanceId"`
	AssociatedInstanceType string                `json:"AssociatedInstanceType" xml:"AssociatedInstanceType"`
	Status                 string                `json:"Status" xml:"Status"`
	NetworkType            string                `json:"NetworkType" xml:"NetworkType"`
	RealBandwidth          int                   `json:"RealBandwidth" xml:"RealBandwidth"`
	AllocationTime         string                `json:"AllocationTime" xml:"AllocationTime"`
	Ipv6Isp                string                `json:"Ipv6Isp" xml:"Ipv6Isp"`
	Ipv6InternetBandwidth  Ipv6InternetBandwidth `json:"Ipv6InternetBandwidth" xml:"Ipv6InternetBandwidth"`
}