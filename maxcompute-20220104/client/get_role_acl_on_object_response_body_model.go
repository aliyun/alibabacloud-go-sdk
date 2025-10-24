// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoleAclOnObjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetRoleAclOnObjectResponseBodyData) *GetRoleAclOnObjectResponseBody
	GetData() *GetRoleAclOnObjectResponseBodyData
	SetRequestId(v string) *GetRoleAclOnObjectResponseBody
	GetRequestId() *string
}

type GetRoleAclOnObjectResponseBody struct {
	// The returned data
	Data *GetRoleAclOnObjectResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0bc1366d16686529650188023ef87f
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetRoleAclOnObjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRoleAclOnObjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetRoleAclOnObjectResponseBody) GetData() *GetRoleAclOnObjectResponseBodyData {
	return s.Data
}

func (s *GetRoleAclOnObjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRoleAclOnObjectResponseBody) SetData(v *GetRoleAclOnObjectResponseBodyData) *GetRoleAclOnObjectResponseBody {
	s.Data = v
	return s
}

func (s *GetRoleAclOnObjectResponseBody) SetRequestId(v string) *GetRoleAclOnObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRoleAclOnObjectResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRoleAclOnObjectResponseBodyData struct {
	// The operations that were performed on the object.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
}

func (s GetRoleAclOnObjectResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRoleAclOnObjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRoleAclOnObjectResponseBodyData) GetActions() []*string {
	return s.Actions
}

func (s *GetRoleAclOnObjectResponseBodyData) SetActions(v []*string) *GetRoleAclOnObjectResponseBodyData {
	s.Actions = v
	return s
}

func (s *GetRoleAclOnObjectResponseBodyData) Validate() error {
	return dara.Validate(s)
}
