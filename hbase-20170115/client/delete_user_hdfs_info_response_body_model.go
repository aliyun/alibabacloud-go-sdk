// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserHdfsInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteUserHdfsInfoResponseBody
	GetClusterId() *string
	SetRequestId(v string) *DeleteUserHdfsInfoResponseBody
	GetRequestId() *string
}

type DeleteUserHdfsInfoResponseBody struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserHdfsInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserHdfsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserHdfsInfoResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteUserHdfsInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserHdfsInfoResponseBody) SetClusterId(v string) *DeleteUserHdfsInfoResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DeleteUserHdfsInfoResponseBody) SetRequestId(v string) *DeleteUserHdfsInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserHdfsInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
