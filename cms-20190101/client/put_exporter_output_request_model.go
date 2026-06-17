// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutExporterOutputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigJson(v string) *PutExporterOutputRequest
	GetConfigJson() *string
	SetDesc(v string) *PutExporterOutputRequest
	GetDesc() *string
	SetDestName(v string) *PutExporterOutputRequest
	GetDestName() *string
	SetDestType(v string) *PutExporterOutputRequest
	GetDestType() *string
	SetRegionId(v string) *PutExporterOutputRequest
	GetRegionId() *string
}

type PutExporterOutputRequest struct {
	// The configuration of the data export. The value is a JSONObject string that must contain the following fields:
	//
	// - endpoint: the domain name that corresponds to the data of Log Service (SLS).
	//
	// - project: the project.
	//
	// - logstore: the Logstore.
	//
	// - ak: the AccessKey ID.
	//
	// - as: the AccessKey secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// { "endpoint": "http://cn-qingdao-share.log.aliyuncs.com", "project": "exporter", "logstore": "exporter","ak": "LTAIp*******", "userId": "17754********", "as": "TxHwuJ8yAb3AU******"}
	ConfigJson *string `json:"ConfigJson,omitempty" xml:"ConfigJson,omitempty"`
	// The description of the configuration.
	//
	// example:
	//
	// CPU metric export
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The name of the configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// exporterConfig
	DestName *string `json:"DestName,omitempty" xml:"DestName,omitempty"`
	// The product to which the data is exported.
	//
	// example:
	//
	// sls
	DestType *string `json:"DestType,omitempty" xml:"DestType,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s PutExporterOutputRequest) String() string {
	return dara.Prettify(s)
}

func (s PutExporterOutputRequest) GoString() string {
	return s.String()
}

func (s *PutExporterOutputRequest) GetConfigJson() *string {
	return s.ConfigJson
}

func (s *PutExporterOutputRequest) GetDesc() *string {
	return s.Desc
}

func (s *PutExporterOutputRequest) GetDestName() *string {
	return s.DestName
}

func (s *PutExporterOutputRequest) GetDestType() *string {
	return s.DestType
}

func (s *PutExporterOutputRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PutExporterOutputRequest) SetConfigJson(v string) *PutExporterOutputRequest {
	s.ConfigJson = &v
	return s
}

func (s *PutExporterOutputRequest) SetDesc(v string) *PutExporterOutputRequest {
	s.Desc = &v
	return s
}

func (s *PutExporterOutputRequest) SetDestName(v string) *PutExporterOutputRequest {
	s.DestName = &v
	return s
}

func (s *PutExporterOutputRequest) SetDestType(v string) *PutExporterOutputRequest {
	s.DestType = &v
	return s
}

func (s *PutExporterOutputRequest) SetRegionId(v string) *PutExporterOutputRequest {
	s.RegionId = &v
	return s
}

func (s *PutExporterOutputRequest) Validate() error {
	return dara.Validate(s)
}
