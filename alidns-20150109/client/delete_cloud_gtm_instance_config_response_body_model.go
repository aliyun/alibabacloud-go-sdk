// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudGtmInstanceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCloudGtmInstanceConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCloudGtmInstanceConfigResponseBody
	GetSuccess() *bool
}

type DeleteCloudGtmInstanceConfigResponseBody struct {
	// Unique request identification code.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation to delete domain instance configurations was successful:
	//
	// - true: Operation successful - false: Operation failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCloudGtmInstanceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudGtmInstanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCloudGtmInstanceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCloudGtmInstanceConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCloudGtmInstanceConfigResponseBody) SetRequestId(v string) *DeleteCloudGtmInstanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCloudGtmInstanceConfigResponseBody) SetSuccess(v bool) *DeleteCloudGtmInstanceConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCloudGtmInstanceConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
