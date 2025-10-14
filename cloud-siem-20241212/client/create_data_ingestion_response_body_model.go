// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataIngestionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataIngestionId(v string) *CreateDataIngestionResponseBody
	GetDataIngestionId() *string
	SetRequestId(v string) *CreateDataIngestionResponseBody
	GetRequestId() *string
}

type CreateDataIngestionResponseBody struct {
	// example:
	//
	// di-yxtm3l2rwa7fr5uvxtc7。
	DataIngestionId *string `json:"DataIngestionId,omitempty" xml:"DataIngestionId,omitempty"`
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataIngestionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataIngestionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataIngestionResponseBody) GetDataIngestionId() *string {
	return s.DataIngestionId
}

func (s *CreateDataIngestionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataIngestionResponseBody) SetDataIngestionId(v string) *CreateDataIngestionResponseBody {
	s.DataIngestionId = &v
	return s
}

func (s *CreateDataIngestionResponseBody) SetRequestId(v string) *CreateDataIngestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataIngestionResponseBody) Validate() error {
	return dara.Validate(s)
}
