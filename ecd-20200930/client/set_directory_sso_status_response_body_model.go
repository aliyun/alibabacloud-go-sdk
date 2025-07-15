// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDirectorySsoStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDirectorySsoStatusResponseBody
	GetRequestId() *string
}

type SetDirectorySsoStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDirectorySsoStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDirectorySsoStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetDirectorySsoStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDirectorySsoStatusResponseBody) SetRequestId(v string) *SetDirectorySsoStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDirectorySsoStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
