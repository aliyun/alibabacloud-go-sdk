// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRecordResponseBody
	GetRequestId() *string
}

type DeleteRecordResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F61CDR30-E83C-4FDA-BF73-9A94CDD44229
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecordResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRecordResponseBody) SetRequestId(v string) *DeleteRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
