// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBaselineCheckWhiteRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteBaselineCheckWhiteRecordResponseBody
	GetRequestId() *string
}

type DeleteBaselineCheckWhiteRecordResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E10BAF1C-A6C5-51E2-866C-76D5922E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBaselineCheckWhiteRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBaselineCheckWhiteRecordResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBaselineCheckWhiteRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBaselineCheckWhiteRecordResponseBody) SetRequestId(v string) *DeleteBaselineCheckWhiteRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBaselineCheckWhiteRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
