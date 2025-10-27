// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetImageSensitiveFileStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SetImageSensitiveFileStatusResponseBody
	GetCode() *string
	SetData(v *SetImageSensitiveFileStatusResponseBodyData) *SetImageSensitiveFileStatusResponseBody
	GetData() *SetImageSensitiveFileStatusResponseBodyData
	SetMessage(v string) *SetImageSensitiveFileStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetImageSensitiveFileStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetImageSensitiveFileStatusResponseBody
	GetSuccess() *bool
}

type SetImageSensitiveFileStatusResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *SetImageSensitiveFileStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20456DD5-5CBF-5015-9173-12CA4246****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetImageSensitiveFileStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetImageSensitiveFileStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetImageSensitiveFileStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *SetImageSensitiveFileStatusResponseBody) GetData() *SetImageSensitiveFileStatusResponseBodyData {
	return s.Data
}

func (s *SetImageSensitiveFileStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetImageSensitiveFileStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetImageSensitiveFileStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetImageSensitiveFileStatusResponseBody) SetCode(v string) *SetImageSensitiveFileStatusResponseBody {
	s.Code = &v
	return s
}

func (s *SetImageSensitiveFileStatusResponseBody) SetData(v *SetImageSensitiveFileStatusResponseBodyData) *SetImageSensitiveFileStatusResponseBody {
	s.Data = v
	return s
}

func (s *SetImageSensitiveFileStatusResponseBody) SetMessage(v string) *SetImageSensitiveFileStatusResponseBody {
	s.Message = &v
	return s
}

func (s *SetImageSensitiveFileStatusResponseBody) SetRequestId(v string) *SetImageSensitiveFileStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetImageSensitiveFileStatusResponseBody) SetSuccess(v bool) *SetImageSensitiveFileStatusResponseBody {
	s.Success = &v
	return s
}

func (s *SetImageSensitiveFileStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetImageSensitiveFileStatusResponseBodyData struct {
	// The primary key ID of the database.
	//
	// example:
	//
	// 18551
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SetImageSensitiveFileStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SetImageSensitiveFileStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *SetImageSensitiveFileStatusResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *SetImageSensitiveFileStatusResponseBodyData) SetId(v int64) *SetImageSensitiveFileStatusResponseBodyData {
	s.Id = &v
	return s
}

func (s *SetImageSensitiveFileStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
