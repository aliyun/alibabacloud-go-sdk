// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMaterializedViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyMaterializedViewResponseBody
	GetRequestId() *string
}

type ModifyMaterializedViewResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// CF8BA882-C9D8-5BEF-AAA3-AD03875F3E18
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyMaterializedViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyMaterializedViewResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMaterializedViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyMaterializedViewResponseBody) SetRequestId(v string) *ModifyMaterializedViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMaterializedViewResponseBody) Validate() error {
	return dara.Validate(s)
}
