// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountZonalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAccountZonalResponseBody
	GetRequestId() *string
}

type CreateAccountZonalResponseBody struct {
	// example:
	//
	// CED079B7-A408-41A1-BFF1-EC608E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccountZonalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountZonalResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccountZonalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAccountZonalResponseBody) SetRequestId(v string) *CreateAccountZonalResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccountZonalResponseBody) Validate() error {
	return dara.Validate(s)
}
