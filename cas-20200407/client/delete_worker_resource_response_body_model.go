// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkerResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWorkerResourceResponseBody
	GetRequestId() *string
}

type DeleteWorkerResourceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EA69E364-5CBB-50E8-BF09-E8CAA396A4F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWorkerResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkerResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkerResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWorkerResourceResponseBody) SetRequestId(v string) *DeleteWorkerResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkerResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
