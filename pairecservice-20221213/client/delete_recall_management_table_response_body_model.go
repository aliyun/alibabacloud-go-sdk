// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecallManagementTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRecallManagementTableResponseBody
	GetRequestId() *string
}

type DeleteRecallManagementTableResponseBody struct {
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRecallManagementTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecallManagementTableResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRecallManagementTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRecallManagementTableResponseBody) SetRequestId(v string) *DeleteRecallManagementTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRecallManagementTableResponseBody) Validate() error {
	return dara.Validate(s)
}
