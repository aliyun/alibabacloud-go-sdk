// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDocResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UploadDocResponseBody
	GetCode() *string
	SetData(v *UploadDocResponseBodyData) *UploadDocResponseBody
	GetData() *UploadDocResponseBodyData
	SetHttpStatusCode(v int32) *UploadDocResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UploadDocResponseBody
	GetMessage() *string
	SetRequestId(v string) *UploadDocResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UploadDocResponseBody
	GetSuccess() *bool
}

type UploadDocResponseBody struct {
	// example:
	//
	// successful
	Code *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UploadDocResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadDocResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadDocResponseBody) GoString() string {
	return s.String()
}

func (s *UploadDocResponseBody) GetCode() *string {
	return s.Code
}

func (s *UploadDocResponseBody) GetData() *UploadDocResponseBodyData {
	return s.Data
}

func (s *UploadDocResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UploadDocResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UploadDocResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadDocResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UploadDocResponseBody) SetCode(v string) *UploadDocResponseBody {
	s.Code = &v
	return s
}

func (s *UploadDocResponseBody) SetData(v *UploadDocResponseBodyData) *UploadDocResponseBody {
	s.Data = v
	return s
}

func (s *UploadDocResponseBody) SetHttpStatusCode(v int32) *UploadDocResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UploadDocResponseBody) SetMessage(v string) *UploadDocResponseBody {
	s.Message = &v
	return s
}

func (s *UploadDocResponseBody) SetRequestId(v string) *UploadDocResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadDocResponseBody) SetSuccess(v bool) *UploadDocResponseBody {
	s.Success = &v
	return s
}

func (s *UploadDocResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UploadDocResponseBodyData struct {
	DocIds     []*string `json:"DocIds,omitempty" xml:"DocIds,omitempty" type:"Repeated"`
	ExistedIds []*string `json:"ExistedIds,omitempty" xml:"ExistedIds,omitempty" type:"Repeated"`
}

func (s UploadDocResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UploadDocResponseBodyData) GoString() string {
	return s.String()
}

func (s *UploadDocResponseBodyData) GetDocIds() []*string {
	return s.DocIds
}

func (s *UploadDocResponseBodyData) GetExistedIds() []*string {
	return s.ExistedIds
}

func (s *UploadDocResponseBodyData) SetDocIds(v []*string) *UploadDocResponseBodyData {
	s.DocIds = v
	return s
}

func (s *UploadDocResponseBodyData) SetExistedIds(v []*string) *UploadDocResponseBodyData {
	s.ExistedIds = v
	return s
}

func (s *UploadDocResponseBodyData) Validate() error {
	return dara.Validate(s)
}
