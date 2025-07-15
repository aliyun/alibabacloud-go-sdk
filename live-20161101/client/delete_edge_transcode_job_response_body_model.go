// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEdgeTranscodeJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEdgeTranscodeJobResponseBody
	GetRequestId() *string
}

type DeleteEdgeTranscodeJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEdgeTranscodeJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEdgeTranscodeJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEdgeTranscodeJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEdgeTranscodeJobResponseBody) SetRequestId(v string) *DeleteEdgeTranscodeJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEdgeTranscodeJobResponseBody) Validate() error {
	return dara.Validate(s)
}
