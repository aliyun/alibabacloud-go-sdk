// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGroupIdByGroupNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *QueryGroupIdByGroupNameResponseBody
	GetGroupId() *int64
	SetRequestId(v string) *QueryGroupIdByGroupNameResponseBody
	GetRequestId() *string
}

type QueryGroupIdByGroupNameResponseBody struct {
	// The ID of the asset group.
	//
	// example:
	//
	// 9935302
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 24A20733-10A0-4AF6-BE6B-E3322413BB68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryGroupIdByGroupNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryGroupIdByGroupNameResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGroupIdByGroupNameResponseBody) GetGroupId() *int64 {
	return s.GroupId
}

func (s *QueryGroupIdByGroupNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryGroupIdByGroupNameResponseBody) SetGroupId(v int64) *QueryGroupIdByGroupNameResponseBody {
	s.GroupId = &v
	return s
}

func (s *QueryGroupIdByGroupNameResponseBody) SetRequestId(v string) *QueryGroupIdByGroupNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryGroupIdByGroupNameResponseBody) Validate() error {
	return dara.Validate(s)
}
