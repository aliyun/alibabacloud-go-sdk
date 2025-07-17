// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeEnvironmentFeatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpgradeEnvironmentFeatureResponseBody
	GetCode() *int32
	SetData(v map[string]*string) *UpgradeEnvironmentFeatureResponseBody
	GetData() map[string]*string
	SetMessage(v string) *UpgradeEnvironmentFeatureResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpgradeEnvironmentFeatureResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpgradeEnvironmentFeatureResponseBody
	GetSuccess() *bool
}

type UpgradeEnvironmentFeatureResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The status of the job.
	Data map[string]*string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 01FF8DD9-A09C-47A1-895A-B6E321BE77B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpgradeEnvironmentFeatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeEnvironmentFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeEnvironmentFeatureResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpgradeEnvironmentFeatureResponseBody) GetData() map[string]*string {
	return s.Data
}

func (s *UpgradeEnvironmentFeatureResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpgradeEnvironmentFeatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeEnvironmentFeatureResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpgradeEnvironmentFeatureResponseBody) SetCode(v int32) *UpgradeEnvironmentFeatureResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradeEnvironmentFeatureResponseBody) SetData(v map[string]*string) *UpgradeEnvironmentFeatureResponseBody {
	s.Data = v
	return s
}

func (s *UpgradeEnvironmentFeatureResponseBody) SetMessage(v string) *UpgradeEnvironmentFeatureResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeEnvironmentFeatureResponseBody) SetRequestId(v string) *UpgradeEnvironmentFeatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeEnvironmentFeatureResponseBody) SetSuccess(v bool) *UpgradeEnvironmentFeatureResponseBody {
	s.Success = &v
	return s
}

func (s *UpgradeEnvironmentFeatureResponseBody) Validate() error {
	return dara.Validate(s)
}
