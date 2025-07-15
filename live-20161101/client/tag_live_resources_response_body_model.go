// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagLiveResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TagLiveResourcesResponseBody
	GetRequestId() *string
}

type TagLiveResourcesResponseBody struct {
	// example:
	//
	// 34AB41F1-04A5-496F-8C8D-634BDBE6A9FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagLiveResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TagLiveResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagLiveResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TagLiveResourcesResponseBody) SetRequestId(v string) *TagLiveResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagLiveResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
