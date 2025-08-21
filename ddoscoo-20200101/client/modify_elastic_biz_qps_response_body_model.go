// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticBizQpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyElasticBizQpsResponseBody
	GetRequestId() *string
}

type ModifyElasticBizQpsResponseBody struct {
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyElasticBizQpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticBizQpsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyElasticBizQpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyElasticBizQpsResponseBody) SetRequestId(v string) *ModifyElasticBizQpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyElasticBizQpsResponseBody) Validate() error {
	return dara.Validate(s)
}
