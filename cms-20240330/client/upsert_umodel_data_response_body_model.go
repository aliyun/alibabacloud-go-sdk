// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertUmodelDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpsertUmodelDataResponseBody
	GetRequestId() *string
}

type UpsertUmodelDataResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 0CEC5375-C554-562B-A65F-9A629907C1F0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpsertUmodelDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpsertUmodelDataResponseBody) GoString() string {
	return s.String()
}

func (s *UpsertUmodelDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpsertUmodelDataResponseBody) SetRequestId(v string) *UpsertUmodelDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpsertUmodelDataResponseBody) Validate() error {
	return dara.Validate(s)
}
