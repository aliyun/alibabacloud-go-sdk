// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyEnableStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetPolicyEnableStatusResponseBody
	GetRequestId() *string
	SetStatusModels(v []*GetPolicyEnableStatusResponseBodyStatusModels) *GetPolicyEnableStatusResponseBody
	GetStatusModels() []*GetPolicyEnableStatusResponseBodyStatusModels
}

type GetPolicyEnableStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6C8DF1B1-C65F-5D3A-9FDA-26A4683BB36B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the Tag Policy feature.
	StatusModels []*GetPolicyEnableStatusResponseBodyStatusModels `json:"StatusModels,omitempty" xml:"StatusModels,omitempty" type:"Repeated"`
}

func (s GetPolicyEnableStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyEnableStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetPolicyEnableStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPolicyEnableStatusResponseBody) GetStatusModels() []*GetPolicyEnableStatusResponseBodyStatusModels {
	return s.StatusModels
}

func (s *GetPolicyEnableStatusResponseBody) SetRequestId(v string) *GetPolicyEnableStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPolicyEnableStatusResponseBody) SetStatusModels(v []*GetPolicyEnableStatusResponseBodyStatusModels) *GetPolicyEnableStatusResponseBody {
	s.StatusModels = v
	return s
}

func (s *GetPolicyEnableStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetPolicyEnableStatusResponseBodyStatusModels struct {
	// The status of the Tag Policy feature. Valid values:
	//
	// 	- PendingEnable: The feature is being enabled.
	//
	// 	- Enabled: The feature is enabled.
	//
	// 	- Closing: The feature is being disabled.
	//
	// 	- Disabled: The feature is disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The mode of the Tag Policy feature. Valid values:
	//
	// 	- USER: single-account mode
	//
	// 	- RD: multi-account mode
	//
	// For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](https://help.aliyun.com/document_detail/417434.html).
	//
	// example:
	//
	// RD
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s GetPolicyEnableStatusResponseBodyStatusModels) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyEnableStatusResponseBodyStatusModels) GoString() string {
	return s.String()
}

func (s *GetPolicyEnableStatusResponseBodyStatusModels) GetStatus() *string {
	return s.Status
}

func (s *GetPolicyEnableStatusResponseBodyStatusModels) GetUserType() *string {
	return s.UserType
}

func (s *GetPolicyEnableStatusResponseBodyStatusModels) SetStatus(v string) *GetPolicyEnableStatusResponseBodyStatusModels {
	s.Status = &v
	return s
}

func (s *GetPolicyEnableStatusResponseBodyStatusModels) SetUserType(v string) *GetPolicyEnableStatusResponseBodyStatusModels {
	s.UserType = &v
	return s
}

func (s *GetPolicyEnableStatusResponseBodyStatusModels) Validate() error {
	return dara.Validate(s)
}
