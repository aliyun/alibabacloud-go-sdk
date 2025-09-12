// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLdpsComputeGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *GetLdpsComputeGroupResponseBody
	GetGroupName() *string
	SetProperties(v map[string]interface{}) *GetLdpsComputeGroupResponseBody
	GetProperties() map[string]interface{}
	SetRequestId(v string) *GetLdpsComputeGroupResponseBody
	GetRequestId() *string
}

type GetLdpsComputeGroupResponseBody struct {
	GroupName  *string                `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Properties map[string]interface{} `json:"Properties,omitempty" xml:"Properties,omitempty"`
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLdpsComputeGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLdpsComputeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetLdpsComputeGroupResponseBody) GetGroupName() *string {
	return s.GroupName
}

func (s *GetLdpsComputeGroupResponseBody) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *GetLdpsComputeGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLdpsComputeGroupResponseBody) SetGroupName(v string) *GetLdpsComputeGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *GetLdpsComputeGroupResponseBody) SetProperties(v map[string]interface{}) *GetLdpsComputeGroupResponseBody {
	s.Properties = v
	return s
}

func (s *GetLdpsComputeGroupResponseBody) SetRequestId(v string) *GetLdpsComputeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLdpsComputeGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
