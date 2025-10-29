// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetColumnResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetColumn(v *Column) *GetColumnResponseBody
	GetColumn() *Column
	SetRequestId(v string) *GetColumnResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetColumnResponseBody
	GetSuccess() *bool
}

type GetColumnResponseBody struct {
	// The columns in the table.
	Column *Column `json:"Column,omitempty" xml:"Column,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D1E2E5BC-xxxx-xxxx-xxxx-xxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetColumnResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetColumnResponseBody) GoString() string {
	return s.String()
}

func (s *GetColumnResponseBody) GetColumn() *Column {
	return s.Column
}

func (s *GetColumnResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetColumnResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetColumnResponseBody) SetColumn(v *Column) *GetColumnResponseBody {
	s.Column = v
	return s
}

func (s *GetColumnResponseBody) SetRequestId(v string) *GetColumnResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetColumnResponseBody) SetSuccess(v bool) *GetColumnResponseBody {
	s.Success = &v
	return s
}

func (s *GetColumnResponseBody) Validate() error {
	if s.Column != nil {
		if err := s.Column.Validate(); err != nil {
			return err
		}
	}
	return nil
}
