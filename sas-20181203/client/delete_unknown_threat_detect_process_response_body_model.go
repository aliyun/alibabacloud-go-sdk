// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUnknownThreatDetectProcessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUnknownThreatDetectProcessResponseBody
	GetRequestId() *string
}

type DeleteUnknownThreatDetectProcessResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 3956048F-9D73-5EDB-834B-4827BB******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUnknownThreatDetectProcessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUnknownThreatDetectProcessResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUnknownThreatDetectProcessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUnknownThreatDetectProcessResponseBody) SetRequestId(v string) *DeleteUnknownThreatDetectProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUnknownThreatDetectProcessResponseBody) Validate() error {
	return dara.Validate(s)
}
