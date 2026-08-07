// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishAppConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *PublishAppConfigResponseBody
	GetData() *bool
	SetRequestId(v string) *PublishAppConfigResponseBody
	GetRequestId() *string
}

type PublishAppConfigResponseBody struct {
	// The returned data.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID assigned by the backend to uniquely identify a request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishAppConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishAppConfigResponseBody) GoString() string {
	return s.String()
}

func (s *PublishAppConfigResponseBody) GetData() *bool {
	return s.Data
}

func (s *PublishAppConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishAppConfigResponseBody) SetData(v bool) *PublishAppConfigResponseBody {
	s.Data = &v
	return s
}

func (s *PublishAppConfigResponseBody) SetRequestId(v string) *PublishAppConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishAppConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
