// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLibrarySchemaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetLibrarySchemaResponseBody
	GetRequestId() *string
	SetResult(v *LibrarySchema) *GetLibrarySchemaResponseBody
	GetResult() *LibrarySchema
}

type GetLibrarySchemaResponseBody struct {
	RequestId *string        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *LibrarySchema `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetLibrarySchemaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLibrarySchemaResponseBody) GoString() string {
	return s.String()
}

func (s *GetLibrarySchemaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLibrarySchemaResponseBody) GetResult() *LibrarySchema {
	return s.Result
}

func (s *GetLibrarySchemaResponseBody) SetRequestId(v string) *GetLibrarySchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLibrarySchemaResponseBody) SetResult(v *LibrarySchema) *GetLibrarySchemaResponseBody {
	s.Result = v
	return s
}

func (s *GetLibrarySchemaResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}
