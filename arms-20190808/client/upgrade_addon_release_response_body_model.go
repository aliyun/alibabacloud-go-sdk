// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeAddonReleaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpgradeAddonReleaseResponseBody
	GetCode() *int32
	SetData(v string) *UpgradeAddonReleaseResponseBody
	GetData() *string
	SetMessage(v string) *UpgradeAddonReleaseResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpgradeAddonReleaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpgradeAddonReleaseResponseBody
	GetSuccess() *bool
}

type UpgradeAddonReleaseResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message returned if the request parameters are invalid.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 27E653FA-5958-45BE-8AA9-14D884DC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpgradeAddonReleaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeAddonReleaseResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeAddonReleaseResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpgradeAddonReleaseResponseBody) GetData() *string {
	return s.Data
}

func (s *UpgradeAddonReleaseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpgradeAddonReleaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeAddonReleaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpgradeAddonReleaseResponseBody) SetCode(v int32) *UpgradeAddonReleaseResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradeAddonReleaseResponseBody) SetData(v string) *UpgradeAddonReleaseResponseBody {
	s.Data = &v
	return s
}

func (s *UpgradeAddonReleaseResponseBody) SetMessage(v string) *UpgradeAddonReleaseResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeAddonReleaseResponseBody) SetRequestId(v string) *UpgradeAddonReleaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeAddonReleaseResponseBody) SetSuccess(v bool) *UpgradeAddonReleaseResponseBody {
	s.Success = &v
	return s
}

func (s *UpgradeAddonReleaseResponseBody) Validate() error {
	return dara.Validate(s)
}
