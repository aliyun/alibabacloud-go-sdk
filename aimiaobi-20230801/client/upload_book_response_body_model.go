// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadBookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UploadBookResponseBody
	GetCode() *string
	SetData(v *UploadBookResponseBodyData) *UploadBookResponseBody
	GetData() *UploadBookResponseBodyData
	SetHttpStatusCode(v int32) *UploadBookResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UploadBookResponseBody
	GetMessage() *string
	SetRequestId(v string) *UploadBookResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UploadBookResponseBody
	GetSuccess() *bool
}

type UploadBookResponseBody struct {
	// example:
	//
	// successful
	Code *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UploadBookResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadBookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadBookResponseBody) GoString() string {
	return s.String()
}

func (s *UploadBookResponseBody) GetCode() *string {
	return s.Code
}

func (s *UploadBookResponseBody) GetData() *UploadBookResponseBodyData {
	return s.Data
}

func (s *UploadBookResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UploadBookResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UploadBookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadBookResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UploadBookResponseBody) SetCode(v string) *UploadBookResponseBody {
	s.Code = &v
	return s
}

func (s *UploadBookResponseBody) SetData(v *UploadBookResponseBodyData) *UploadBookResponseBody {
	s.Data = v
	return s
}

func (s *UploadBookResponseBody) SetHttpStatusCode(v int32) *UploadBookResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UploadBookResponseBody) SetMessage(v string) *UploadBookResponseBody {
	s.Message = &v
	return s
}

func (s *UploadBookResponseBody) SetRequestId(v string) *UploadBookResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadBookResponseBody) SetSuccess(v bool) *UploadBookResponseBody {
	s.Success = &v
	return s
}

func (s *UploadBookResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UploadBookResponseBodyData struct {
	DocIds     []*string `json:"DocIds,omitempty" xml:"DocIds,omitempty" type:"Repeated"`
	ExistedIds []*string `json:"ExistedIds,omitempty" xml:"ExistedIds,omitempty" type:"Repeated"`
}

func (s UploadBookResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UploadBookResponseBodyData) GoString() string {
	return s.String()
}

func (s *UploadBookResponseBodyData) GetDocIds() []*string {
	return s.DocIds
}

func (s *UploadBookResponseBodyData) GetExistedIds() []*string {
	return s.ExistedIds
}

func (s *UploadBookResponseBodyData) SetDocIds(v []*string) *UploadBookResponseBodyData {
	s.DocIds = v
	return s
}

func (s *UploadBookResponseBodyData) SetExistedIds(v []*string) *UploadBookResponseBodyData {
	s.ExistedIds = v
	return s
}

func (s *UploadBookResponseBodyData) Validate() error {
	return dara.Validate(s)
}
