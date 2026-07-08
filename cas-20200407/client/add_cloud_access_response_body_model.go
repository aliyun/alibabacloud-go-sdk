// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCloudAccessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddCloudAccessResponseBody
	GetRequestId() *string
}

type AddCloudAccessResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 285BBE08-F12B-5A04-97BC-09EA7FF18646
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddCloudAccessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCloudAccessResponseBody) GoString() string {
	return s.String()
}

func (s *AddCloudAccessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCloudAccessResponseBody) SetRequestId(v string) *AddCloudAccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCloudAccessResponseBody) Validate() error {
	return dara.Validate(s)
}
