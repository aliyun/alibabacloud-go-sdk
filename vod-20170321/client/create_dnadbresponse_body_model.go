// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDNADBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBId(v string) *CreateDNADBResponseBody
	GetDBId() *string
	SetRequestId(v string) *CreateDNADBResponseBody
	GetRequestId() *string
}

type CreateDNADBResponseBody struct {
	DBId      *string `json:"DBId,omitempty" xml:"DBId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDNADBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDNADBResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDNADBResponseBody) GetDBId() *string {
	return s.DBId
}

func (s *CreateDNADBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDNADBResponseBody) SetDBId(v string) *CreateDNADBResponseBody {
	s.DBId = &v
	return s
}

func (s *CreateDNADBResponseBody) SetRequestId(v string) *CreateDNADBResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDNADBResponseBody) Validate() error {
	return dara.Validate(s)
}
