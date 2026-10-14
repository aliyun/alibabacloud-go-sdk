// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigNetworkRegionBlockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *ConfigNetworkRegionBlockRequest
	GetConfig() *string
	SetInstanceId(v string) *ConfigNetworkRegionBlockRequest
	GetInstanceId() *string
}

type ConfigNetworkRegionBlockRequest struct {
	// The configuration of the location blacklist, expressed as a JSON string with the following structure:
	//
	// - **RegionBlockSwitch**: String type. Required. The switch status of the location blacklist. Valid values:
	//
	//     - **on**: enabled
	//
	//     - **off**: disabled
	//
	// - **Countries**: Array type. Optional. The list of international region codes to block.
	//
	//     > For more information about international region codes, see [Parameters for China and international regions](https://help.aliyun.com/document_detail/167926.html).
	//
	//
	// - **Provinces**: Array type. Optional. The list of China region codes to block.
	//
	//     > For more information about China region codes, see [Parameters for China and international regions](https://help.aliyun.com/document_detail/167926.html).
	//
	//     For example, `[11,12]` indicates Beijing and Tianjin.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"RegionBlockSwitch":"on","Countries":[],"Provinces":[11,12,13,14,15,21,22,23,31,32,33,34,35,36,37,41,42,43,44,45,46,50,51,52,53,54,61,62,63,64,65,71,81,82]}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The ID of the Anti-DDoS Pro or Anti-DDoS Premium instance.
	//
	// > You can call [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) to query the IDs of all Anti-DDoS Pro or Anti-DDoS Premium instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ConfigNetworkRegionBlockRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigNetworkRegionBlockRequest) GoString() string {
	return s.String()
}

func (s *ConfigNetworkRegionBlockRequest) GetConfig() *string {
	return s.Config
}

func (s *ConfigNetworkRegionBlockRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConfigNetworkRegionBlockRequest) SetConfig(v string) *ConfigNetworkRegionBlockRequest {
	s.Config = &v
	return s
}

func (s *ConfigNetworkRegionBlockRequest) SetInstanceId(v string) *ConfigNetworkRegionBlockRequest {
	s.InstanceId = &v
	return s
}

func (s *ConfigNetworkRegionBlockRequest) Validate() error {
	return dara.Validate(s)
}
