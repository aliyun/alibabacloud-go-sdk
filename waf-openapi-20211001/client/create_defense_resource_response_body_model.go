// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDefenseResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDefenseResourceResponseBody
	GetRequestId() *string
	SetResource(v string) *CreateDefenseResourceResponseBody
	GetResource() *string
}

type CreateDefenseResourceResponseBody struct {
	// example:
	//
	// 1738C613-D054-5191-888B-DC0CF4C3A4A0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// cdX.XXX-call.cn-alb
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
}

func (s CreateDefenseResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDefenseResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDefenseResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDefenseResourceResponseBody) GetResource() *string {
	return s.Resource
}

func (s *CreateDefenseResourceResponseBody) SetRequestId(v string) *CreateDefenseResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDefenseResourceResponseBody) SetResource(v string) *CreateDefenseResourceResponseBody {
	s.Resource = &v
	return s
}

func (s *CreateDefenseResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
