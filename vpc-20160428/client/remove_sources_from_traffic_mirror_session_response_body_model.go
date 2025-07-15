// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSourcesFromTrafficMirrorSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveSourcesFromTrafficMirrorSessionResponseBody
	GetRequestId() *string
}

type RemoveSourcesFromTrafficMirrorSessionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A33B2C6A-89D1-4DEA-A807-A6E8CC552484
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveSourcesFromTrafficMirrorSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveSourcesFromTrafficMirrorSessionResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveSourcesFromTrafficMirrorSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveSourcesFromTrafficMirrorSessionResponseBody) SetRequestId(v string) *RemoveSourcesFromTrafficMirrorSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveSourcesFromTrafficMirrorSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
