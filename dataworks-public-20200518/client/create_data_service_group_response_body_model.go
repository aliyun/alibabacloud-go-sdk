// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataServiceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *CreateDataServiceGroupResponseBody
	GetGroupId() *string
	SetRequestId(v string) *CreateDataServiceGroupResponseBody
	GetRequestId() *string
}

type CreateDataServiceGroupResponseBody struct {
	// The business process ID.
	//
	// example:
	//
	// ds_12345
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataServiceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataServiceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataServiceGroupResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateDataServiceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataServiceGroupResponseBody) SetGroupId(v string) *CreateDataServiceGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *CreateDataServiceGroupResponseBody) SetRequestId(v string) *CreateDataServiceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataServiceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
