// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDirectoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDirectoriesResponseBody
	GetRequestId() *string
}

type DeleteDirectoriesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5CC5E450-FC43-4F5B-B540-9964BD313427
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDirectoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDirectoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDirectoriesResponseBody) SetRequestId(v string) *DeleteDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDirectoriesResponseBody) Validate() error {
	return dara.Validate(s)
}
