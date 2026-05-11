// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveRecordingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveRecordingResponseBody
	GetRequestId() *string
}

type SaveRecordingResponseBody struct {
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveRecordingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveRecordingResponseBody) GoString() string {
	return s.String()
}

func (s *SaveRecordingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveRecordingResponseBody) SetRequestId(v string) *SaveRecordingResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveRecordingResponseBody) Validate() error {
	return dara.Validate(s)
}
