// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetAsyncResultResponseBody
	GetData() *string
	SetRequestId(v string) *GetAsyncResultResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetAsyncResultResponseBody
	GetStatus() *string
}

type GetAsyncResultResponseBody struct {
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// F79E7305-5314-5069-A701-9591AD051902
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAsyncResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncResultResponseBody) GetData() *string {
	return s.Data
}

func (s *GetAsyncResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAsyncResultResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetAsyncResultResponseBody) SetData(v string) *GetAsyncResultResponseBody {
	s.Data = &v
	return s
}

func (s *GetAsyncResultResponseBody) SetRequestId(v string) *GetAsyncResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAsyncResultResponseBody) SetStatus(v string) *GetAsyncResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetAsyncResultResponseBody) Validate() error {
	return dara.Validate(s)
}
