// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCustomerListResponseBody
	GetCode() *string
	SetData(v *GetCustomerListResponseBodyData) *GetCustomerListResponseBody
	GetData() *GetCustomerListResponseBodyData
	SetMessage(v string) *GetCustomerListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCustomerListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCustomerListResponseBody
	GetSuccess() *bool
}

type GetCustomerListResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetCustomerListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- **true**: The call is successful.
	//
	// 	- **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCustomerListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerListResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomerListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCustomerListResponseBody) GetData() *GetCustomerListResponseBodyData {
	return s.Data
}

func (s *GetCustomerListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCustomerListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCustomerListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCustomerListResponseBody) SetCode(v string) *GetCustomerListResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomerListResponseBody) SetData(v *GetCustomerListResponseBodyData) *GetCustomerListResponseBody {
	s.Data = v
	return s
}

func (s *GetCustomerListResponseBody) SetMessage(v string) *GetCustomerListResponseBody {
	s.Message = &v
	return s
}

func (s *GetCustomerListResponseBody) SetRequestId(v string) *GetCustomerListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomerListResponseBody) SetSuccess(v bool) *GetCustomerListResponseBody {
	s.Success = &v
	return s
}

func (s *GetCustomerListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCustomerListResponseBodyData struct {
	// The list of customer IDs.
	UidList []*string `json:"UidList,omitempty" xml:"UidList,omitempty" type:"Repeated"`
}

func (s GetCustomerListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCustomerListResponseBodyData) GetUidList() []*string {
	return s.UidList
}

func (s *GetCustomerListResponseBodyData) SetUidList(v []*string) *GetCustomerListResponseBodyData {
	s.UidList = v
	return s
}

func (s *GetCustomerListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
