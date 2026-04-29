// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMOQuotaAlertThresholdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMOQuotaAlertThresholdResponseBody
	GetRequestId() *string
	SetResults(v []*UpdateMOQuotaAlertThresholdResponseBodyResults) *UpdateMOQuotaAlertThresholdResponseBody
	GetResults() []*UpdateMOQuotaAlertThresholdResponseBodyResults
	SetSuccess(v bool) *UpdateMOQuotaAlertThresholdResponseBody
	GetSuccess() *bool
}

type UpdateMOQuotaAlertThresholdResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*UpdateMOQuotaAlertThresholdResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMOQuotaAlertThresholdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMOQuotaAlertThresholdResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMOQuotaAlertThresholdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMOQuotaAlertThresholdResponseBody) GetResults() []*UpdateMOQuotaAlertThresholdResponseBodyResults {
	return s.Results
}

func (s *UpdateMOQuotaAlertThresholdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMOQuotaAlertThresholdResponseBody) SetRequestId(v string) *UpdateMOQuotaAlertThresholdResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMOQuotaAlertThresholdResponseBody) SetResults(v []*UpdateMOQuotaAlertThresholdResponseBodyResults) *UpdateMOQuotaAlertThresholdResponseBody {
	s.Results = v
	return s
}

func (s *UpdateMOQuotaAlertThresholdResponseBody) SetSuccess(v bool) *UpdateMOQuotaAlertThresholdResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMOQuotaAlertThresholdResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateMOQuotaAlertThresholdResponseBodyResults struct {
	// API Key
	//
	// example:
	//
	// sk-rds-*****
	Apikey *string `json:"Apikey,omitempty" xml:"Apikey,omitempty"`
	// example:
	//
	// rds_copilot***_public_cn-o*****1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// api-*****
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// example:
	//
	// system / custom
	KeyType          *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	ThresholdPercent *int32  `json:"ThresholdPercent,omitempty" xml:"ThresholdPercent,omitempty"`
}

func (s UpdateMOQuotaAlertThresholdResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s UpdateMOQuotaAlertThresholdResponseBodyResults) GoString() string {
	return s.String()
}

func (s *UpdateMOQuotaAlertThresholdResponseBodyResults) GetApikey() *string {
	return s.Apikey
}

func (s *UpdateMOQuotaAlertThresholdResponseBodyResults) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateMOQuotaAlertThresholdResponseBodyResults) GetKeyName() *string {
	return s.KeyName
}

func (s *UpdateMOQuotaAlertThresholdResponseBodyResults) GetKeyType() *string {
	return s.KeyType
}

func (s *UpdateMOQuotaAlertThresholdResponseBodyResults) GetThresholdPercent() *int32 {
	return s.ThresholdPercent
}

func (s *UpdateMOQuotaAlertThresholdResponseBodyResults) SetApikey(v string) *UpdateMOQuotaAlertThresholdResponseBodyResults {
	s.Apikey = &v
	return s
}

func (s *UpdateMOQuotaAlertThresholdResponseBodyResults) SetInstanceId(v string) *UpdateMOQuotaAlertThresholdResponseBodyResults {
	s.InstanceId = &v
	return s
}

func (s *UpdateMOQuotaAlertThresholdResponseBodyResults) SetKeyName(v string) *UpdateMOQuotaAlertThresholdResponseBodyResults {
	s.KeyName = &v
	return s
}

func (s *UpdateMOQuotaAlertThresholdResponseBodyResults) SetKeyType(v string) *UpdateMOQuotaAlertThresholdResponseBodyResults {
	s.KeyType = &v
	return s
}

func (s *UpdateMOQuotaAlertThresholdResponseBodyResults) SetThresholdPercent(v int32) *UpdateMOQuotaAlertThresholdResponseBodyResults {
	s.ThresholdPercent = &v
	return s
}

func (s *UpdateMOQuotaAlertThresholdResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
