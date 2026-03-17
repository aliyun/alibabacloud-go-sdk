// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeSagInstanceFromCcnResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevokeSagInstanceFromCcnResponseBody
	GetRequestId() *string
}

type RevokeSagInstanceFromCcnResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// BCB97739-0CB5-4C94-9A5C-13051FFAB0E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeSagInstanceFromCcnResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeSagInstanceFromCcnResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeSagInstanceFromCcnResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeSagInstanceFromCcnResponseBody) SetRequestId(v string) *RevokeSagInstanceFromCcnResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeSagInstanceFromCcnResponseBody) Validate() error {
	return dara.Validate(s)
}
