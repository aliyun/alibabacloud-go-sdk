// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFolderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteFolderResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteFolderResponseBody
	GetRequestId() *string
}

type DeleteFolderResponseBody struct {
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteFolderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFolderResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFolderResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteFolderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFolderResponseBody) SetData(v bool) *DeleteFolderResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteFolderResponseBody) SetRequestId(v string) *DeleteFolderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFolderResponseBody) Validate() error {
	return dara.Validate(s)
}
