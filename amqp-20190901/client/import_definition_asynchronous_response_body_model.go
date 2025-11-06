// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportDefinitionAsynchronousResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ImportDefinitionAsynchronousResponseBody
	GetCode() *int32
	SetData(v *ImportDefinitionAsynchronousResponseBodyData) *ImportDefinitionAsynchronousResponseBody
	GetData() *ImportDefinitionAsynchronousResponseBodyData
	SetMessage(v string) *ImportDefinitionAsynchronousResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImportDefinitionAsynchronousResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImportDefinitionAsynchronousResponseBody
	GetSuccess() *bool
}

type ImportDefinitionAsynchronousResponseBody struct {
	Code      *int32                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ImportDefinitionAsynchronousResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImportDefinitionAsynchronousResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportDefinitionAsynchronousResponseBody) GoString() string {
	return s.String()
}

func (s *ImportDefinitionAsynchronousResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ImportDefinitionAsynchronousResponseBody) GetData() *ImportDefinitionAsynchronousResponseBodyData {
	return s.Data
}

func (s *ImportDefinitionAsynchronousResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportDefinitionAsynchronousResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportDefinitionAsynchronousResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImportDefinitionAsynchronousResponseBody) SetCode(v int32) *ImportDefinitionAsynchronousResponseBody {
	s.Code = &v
	return s
}

func (s *ImportDefinitionAsynchronousResponseBody) SetData(v *ImportDefinitionAsynchronousResponseBodyData) *ImportDefinitionAsynchronousResponseBody {
	s.Data = v
	return s
}

func (s *ImportDefinitionAsynchronousResponseBody) SetMessage(v string) *ImportDefinitionAsynchronousResponseBody {
	s.Message = &v
	return s
}

func (s *ImportDefinitionAsynchronousResponseBody) SetRequestId(v string) *ImportDefinitionAsynchronousResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportDefinitionAsynchronousResponseBody) SetSuccess(v bool) *ImportDefinitionAsynchronousResponseBody {
	s.Success = &v
	return s
}

func (s *ImportDefinitionAsynchronousResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImportDefinitionAsynchronousResponseBodyData struct {
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ImportDefinitionAsynchronousResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImportDefinitionAsynchronousResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImportDefinitionAsynchronousResponseBodyData) GetResult() *bool {
	return s.Result
}

func (s *ImportDefinitionAsynchronousResponseBodyData) SetResult(v bool) *ImportDefinitionAsynchronousResponseBodyData {
	s.Result = &v
	return s
}

func (s *ImportDefinitionAsynchronousResponseBodyData) Validate() error {
	return dara.Validate(s)
}
