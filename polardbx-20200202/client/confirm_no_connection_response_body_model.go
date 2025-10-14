// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmNoConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ConfirmNoConnectionResponseBodyData) *ConfirmNoConnectionResponseBody
	GetData() *ConfirmNoConnectionResponseBodyData
	SetMessage(v string) *ConfirmNoConnectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *ConfirmNoConnectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ConfirmNoConnectionResponseBody
	GetSuccess() *bool
}

type ConfirmNoConnectionResponseBody struct {
	Data *ConfirmNoConnectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfirmNoConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfirmNoConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmNoConnectionResponseBody) GetData() *ConfirmNoConnectionResponseBodyData {
	return s.Data
}

func (s *ConfirmNoConnectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ConfirmNoConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfirmNoConnectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ConfirmNoConnectionResponseBody) SetData(v *ConfirmNoConnectionResponseBodyData) *ConfirmNoConnectionResponseBody {
	s.Data = v
	return s
}

func (s *ConfirmNoConnectionResponseBody) SetMessage(v string) *ConfirmNoConnectionResponseBody {
	s.Message = &v
	return s
}

func (s *ConfirmNoConnectionResponseBody) SetRequestId(v string) *ConfirmNoConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfirmNoConnectionResponseBody) SetSuccess(v bool) *ConfirmNoConnectionResponseBody {
	s.Success = &v
	return s
}

func (s *ConfirmNoConnectionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConfirmNoConnectionResponseBodyData struct {
	// example:
	//
	// etx-szr2rr6i*****
	SlinkTaskId *string `json:"SlinkTaskId,omitempty" xml:"SlinkTaskId,omitempty"`
}

func (s ConfirmNoConnectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ConfirmNoConnectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *ConfirmNoConnectionResponseBodyData) GetSlinkTaskId() *string {
	return s.SlinkTaskId
}

func (s *ConfirmNoConnectionResponseBodyData) SetSlinkTaskId(v string) *ConfirmNoConnectionResponseBodyData {
	s.SlinkTaskId = &v
	return s
}

func (s *ConfirmNoConnectionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
