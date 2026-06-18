// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMakeDoubleCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MakeDoubleCallResponseBody
	GetCode() *string
	SetData(v *MakeDoubleCallResponseBodyData) *MakeDoubleCallResponseBody
	GetData() *MakeDoubleCallResponseBodyData
	SetMessage(v string) *MakeDoubleCallResponseBody
	GetMessage() *string
	SetRequestId(v string) *MakeDoubleCallResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MakeDoubleCallResponseBody
	GetSuccess() *bool
}

type MakeDoubleCallResponseBody struct {
	// Status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data.
	Data *MakeDoubleCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Description of the status code.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MakeDoubleCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MakeDoubleCallResponseBody) GoString() string {
	return s.String()
}

func (s *MakeDoubleCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *MakeDoubleCallResponseBody) GetData() *MakeDoubleCallResponseBodyData {
	return s.Data
}

func (s *MakeDoubleCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MakeDoubleCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MakeDoubleCallResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MakeDoubleCallResponseBody) SetCode(v string) *MakeDoubleCallResponseBody {
	s.Code = &v
	return s
}

func (s *MakeDoubleCallResponseBody) SetData(v *MakeDoubleCallResponseBodyData) *MakeDoubleCallResponseBody {
	s.Data = v
	return s
}

func (s *MakeDoubleCallResponseBody) SetMessage(v string) *MakeDoubleCallResponseBody {
	s.Message = &v
	return s
}

func (s *MakeDoubleCallResponseBody) SetRequestId(v string) *MakeDoubleCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *MakeDoubleCallResponseBody) SetSuccess(v bool) *MakeDoubleCallResponseBody {
	s.Success = &v
	return s
}

func (s *MakeDoubleCallResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MakeDoubleCallResponseBodyData struct {
	// Session ID.
	//
	// example:
	//
	// 68255155365620598
	Acid *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
}

func (s MakeDoubleCallResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MakeDoubleCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *MakeDoubleCallResponseBodyData) GetAcid() *string {
	return s.Acid
}

func (s *MakeDoubleCallResponseBodyData) SetAcid(v string) *MakeDoubleCallResponseBodyData {
	s.Acid = &v
	return s
}

func (s *MakeDoubleCallResponseBodyData) Validate() error {
	return dara.Validate(s)
}
