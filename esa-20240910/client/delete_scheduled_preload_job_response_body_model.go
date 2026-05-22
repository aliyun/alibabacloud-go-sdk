// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScheduledPreloadJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteScheduledPreloadJobResponseBody
	GetRequestId() *string
}

type DeleteScheduledPreloadJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C370DAF1-C838-4288-A1A0-9A87633D248E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteScheduledPreloadJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduledPreloadJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScheduledPreloadJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteScheduledPreloadJobResponseBody) SetRequestId(v string) *DeleteScheduledPreloadJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteScheduledPreloadJobResponseBody) Validate() error {
	return dara.Validate(s)
}
