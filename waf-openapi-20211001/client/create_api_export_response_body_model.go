// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiExportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateApiExportResponseBody
	GetRequestId() *string
}

type CreateApiExportResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 62382992-F9AA-52B2-9147-66B3B9E51D74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApiExportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApiExportResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApiExportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApiExportResponseBody) SetRequestId(v string) *CreateApiExportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApiExportResponseBody) Validate() error {
	return dara.Validate(s)
}
