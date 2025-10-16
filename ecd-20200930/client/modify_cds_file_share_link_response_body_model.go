// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCdsFileShareLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyCdsFileShareLinkResponseBody
	GetCode() *string
	SetData(v *CdsFileShareLinkModel) *ModifyCdsFileShareLinkResponseBody
	GetData() *CdsFileShareLinkModel
	SetMessage(v string) *ModifyCdsFileShareLinkResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyCdsFileShareLinkResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyCdsFileShareLinkResponseBody
	GetSuccess() *bool
}

type ModifyCdsFileShareLinkResponseBody struct {
	// The modification result. The value success indicates that the modification is successful. If the modification failed, an error message is returned.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data information.
	Data *CdsFileShareLinkModel `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message that is returned. This parameter is not returned if the value of Code is success.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F7E4322D-D679-5ACB-A909-490D2F0E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request status.
	//
	// Valid values:
	//
	// 	- true: The request is successful.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false: The request fails.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyCdsFileShareLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCdsFileShareLinkResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCdsFileShareLinkResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyCdsFileShareLinkResponseBody) GetData() *CdsFileShareLinkModel {
	return s.Data
}

func (s *ModifyCdsFileShareLinkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyCdsFileShareLinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCdsFileShareLinkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyCdsFileShareLinkResponseBody) SetCode(v string) *ModifyCdsFileShareLinkResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyCdsFileShareLinkResponseBody) SetData(v *CdsFileShareLinkModel) *ModifyCdsFileShareLinkResponseBody {
	s.Data = v
	return s
}

func (s *ModifyCdsFileShareLinkResponseBody) SetMessage(v string) *ModifyCdsFileShareLinkResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyCdsFileShareLinkResponseBody) SetRequestId(v string) *ModifyCdsFileShareLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCdsFileShareLinkResponseBody) SetSuccess(v bool) *ModifyCdsFileShareLinkResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyCdsFileShareLinkResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
