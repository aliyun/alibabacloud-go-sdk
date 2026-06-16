// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCanTrySasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCanTrySasResponseBody
	GetCode() *string
	SetData(v *GetCanTrySasResponseBodyData) *GetCanTrySasResponseBody
	GetData() *GetCanTrySasResponseBodyData
	SetMessage(v string) *GetCanTrySasResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCanTrySasResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCanTrySasResponseBody
	GetSuccess() *bool
}

type GetCanTrySasResponseBody struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetCanTrySasResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCanTrySasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCanTrySasResponseBody) GoString() string {
	return s.String()
}

func (s *GetCanTrySasResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCanTrySasResponseBody) GetData() *GetCanTrySasResponseBodyData {
	return s.Data
}

func (s *GetCanTrySasResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCanTrySasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCanTrySasResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCanTrySasResponseBody) SetCode(v string) *GetCanTrySasResponseBody {
	s.Code = &v
	return s
}

func (s *GetCanTrySasResponseBody) SetData(v *GetCanTrySasResponseBodyData) *GetCanTrySasResponseBody {
	s.Data = v
	return s
}

func (s *GetCanTrySasResponseBody) SetMessage(v string) *GetCanTrySasResponseBody {
	s.Message = &v
	return s
}

func (s *GetCanTrySasResponseBody) SetRequestId(v string) *GetCanTrySasResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCanTrySasResponseBody) SetSuccess(v bool) *GetCanTrySasResponseBody {
	s.Success = &v
	return s
}

func (s *GetCanTrySasResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCanTrySasResponseBodyData struct {
	Body *GetCanTrySasResponseBodyDataBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
}

func (s GetCanTrySasResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCanTrySasResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCanTrySasResponseBodyData) GetBody() *GetCanTrySasResponseBodyDataBody {
	return s.Body
}

func (s *GetCanTrySasResponseBodyData) SetBody(v *GetCanTrySasResponseBodyDataBody) *GetCanTrySasResponseBodyData {
	s.Body = v
	return s
}

func (s *GetCanTrySasResponseBodyData) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCanTrySasResponseBodyDataBody struct {
	Data      *GetCanTrySasResponseBodyDataBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCanTrySasResponseBodyDataBody) String() string {
	return dara.Prettify(s)
}

func (s GetCanTrySasResponseBodyDataBody) GoString() string {
	return s.String()
}

func (s *GetCanTrySasResponseBodyDataBody) GetData() *GetCanTrySasResponseBodyDataBodyData {
	return s.Data
}

func (s *GetCanTrySasResponseBodyDataBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCanTrySasResponseBodyDataBody) SetData(v *GetCanTrySasResponseBodyDataBodyData) *GetCanTrySasResponseBodyDataBody {
	s.Data = v
	return s
}

func (s *GetCanTrySasResponseBodyDataBody) SetRequestId(v string) *GetCanTrySasResponseBodyDataBody {
	s.RequestId = &v
	return s
}

func (s *GetCanTrySasResponseBodyDataBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCanTrySasResponseBodyDataBodyData struct {
	CanTry         *int32   `json:"CanTry,omitempty" xml:"CanTry,omitempty"`
	CanTryVersions []*int32 `json:"CanTryVersions,omitempty" xml:"CanTryVersions,omitempty" type:"Repeated"`
	TryType        *int32   `json:"TryType,omitempty" xml:"TryType,omitempty"`
}

func (s GetCanTrySasResponseBodyDataBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCanTrySasResponseBodyDataBodyData) GoString() string {
	return s.String()
}

func (s *GetCanTrySasResponseBodyDataBodyData) GetCanTry() *int32 {
	return s.CanTry
}

func (s *GetCanTrySasResponseBodyDataBodyData) GetCanTryVersions() []*int32 {
	return s.CanTryVersions
}

func (s *GetCanTrySasResponseBodyDataBodyData) GetTryType() *int32 {
	return s.TryType
}

func (s *GetCanTrySasResponseBodyDataBodyData) SetCanTry(v int32) *GetCanTrySasResponseBodyDataBodyData {
	s.CanTry = &v
	return s
}

func (s *GetCanTrySasResponseBodyDataBodyData) SetCanTryVersions(v []*int32) *GetCanTrySasResponseBodyDataBodyData {
	s.CanTryVersions = v
	return s
}

func (s *GetCanTrySasResponseBodyDataBodyData) SetTryType(v int32) *GetCanTrySasResponseBodyDataBodyData {
	s.TryType = &v
	return s
}

func (s *GetCanTrySasResponseBodyDataBodyData) Validate() error {
	return dara.Validate(s)
}
