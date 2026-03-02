// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLibrarySchemaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLibrarySchemaResponseBody
	GetRequestId() *string
	SetResult(v *LibrarySchema) *UpdateLibrarySchemaResponseBody
	GetResult() *LibrarySchema
}

type UpdateLibrarySchemaResponseBody struct {
	RequestId *string        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *LibrarySchema `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateLibrarySchemaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLibrarySchemaResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLibrarySchemaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLibrarySchemaResponseBody) GetResult() *LibrarySchema {
	return s.Result
}

func (s *UpdateLibrarySchemaResponseBody) SetRequestId(v string) *UpdateLibrarySchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLibrarySchemaResponseBody) SetResult(v *LibrarySchema) *UpdateLibrarySchemaResponseBody {
	s.Result = v
	return s
}

func (s *UpdateLibrarySchemaResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}
