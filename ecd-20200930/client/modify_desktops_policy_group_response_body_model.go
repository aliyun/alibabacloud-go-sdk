// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopsPolicyGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModifyResults(v []*ModifyDesktopsPolicyGroupResponseBodyModifyResults) *ModifyDesktopsPolicyGroupResponseBody
	GetModifyResults() []*ModifyDesktopsPolicyGroupResponseBodyModifyResults
	SetRequestId(v string) *ModifyDesktopsPolicyGroupResponseBody
	GetRequestId() *string
}

type ModifyDesktopsPolicyGroupResponseBody struct {
	// The request results.
	ModifyResults []*ModifyDesktopsPolicyGroupResponseBodyModifyResults `json:"ModifyResults,omitempty" xml:"ModifyResults,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDesktopsPolicyGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopsPolicyGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDesktopsPolicyGroupResponseBody) GetModifyResults() []*ModifyDesktopsPolicyGroupResponseBodyModifyResults {
	return s.ModifyResults
}

func (s *ModifyDesktopsPolicyGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDesktopsPolicyGroupResponseBody) SetModifyResults(v []*ModifyDesktopsPolicyGroupResponseBodyModifyResults) *ModifyDesktopsPolicyGroupResponseBody {
	s.ModifyResults = v
	return s
}

func (s *ModifyDesktopsPolicyGroupResponseBody) SetRequestId(v string) *ModifyDesktopsPolicyGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDesktopsPolicyGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type ModifyDesktopsPolicyGroupResponseBodyModifyResults struct {
	// The returned message. If the request was successful, `success` is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The cloud computer ID.
	//
	// example:
	//
	// ecd-7w78ozhjcwa3u****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The error message returned if the request failed. This parameter is not returned if the value of Code is success.``
	//
	// example:
	//
	// The specified param DesktopId ecd-ia2zw38bi6cm7***	- is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ModifyDesktopsPolicyGroupResponseBodyModifyResults) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopsPolicyGroupResponseBodyModifyResults) GoString() string {
	return s.String()
}

func (s *ModifyDesktopsPolicyGroupResponseBodyModifyResults) GetCode() *string {
	return s.Code
}

func (s *ModifyDesktopsPolicyGroupResponseBodyModifyResults) GetDesktopId() *string {
	return s.DesktopId
}

func (s *ModifyDesktopsPolicyGroupResponseBodyModifyResults) GetMessage() *string {
	return s.Message
}

func (s *ModifyDesktopsPolicyGroupResponseBodyModifyResults) SetCode(v string) *ModifyDesktopsPolicyGroupResponseBodyModifyResults {
	s.Code = &v
	return s
}

func (s *ModifyDesktopsPolicyGroupResponseBodyModifyResults) SetDesktopId(v string) *ModifyDesktopsPolicyGroupResponseBodyModifyResults {
	s.DesktopId = &v
	return s
}

func (s *ModifyDesktopsPolicyGroupResponseBodyModifyResults) SetMessage(v string) *ModifyDesktopsPolicyGroupResponseBodyModifyResults {
	s.Message = &v
	return s
}

func (s *ModifyDesktopsPolicyGroupResponseBodyModifyResults) Validate() error {
	return dara.Validate(s)
}
