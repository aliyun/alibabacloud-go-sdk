// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOnlineDatabaseTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateOnlineDatabaseTaskResponseBody
	GetRequestId() *string
}

type CreateOnlineDatabaseTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1B2EBD14-36F6-4645-A3F9-DE19D321C18F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOnlineDatabaseTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOnlineDatabaseTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOnlineDatabaseTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOnlineDatabaseTaskResponseBody) SetRequestId(v string) *CreateOnlineDatabaseTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOnlineDatabaseTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
