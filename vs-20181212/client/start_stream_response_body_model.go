// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *StartStreamResponseBody
	GetId() *string
	SetName(v string) *StartStreamResponseBody
	GetName() *string
	SetRequestId(v string) *StartStreamResponseBody
	GetRequestId() *string
}

type StartStreamResponseBody struct {
	// example:
	//
	// 323*****997-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 31000000000000000002
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartStreamResponseBody) GoString() string {
	return s.String()
}

func (s *StartStreamResponseBody) GetId() *string {
	return s.Id
}

func (s *StartStreamResponseBody) GetName() *string {
	return s.Name
}

func (s *StartStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartStreamResponseBody) SetId(v string) *StartStreamResponseBody {
	s.Id = &v
	return s
}

func (s *StartStreamResponseBody) SetName(v string) *StartStreamResponseBody {
	s.Name = &v
	return s
}

func (s *StartStreamResponseBody) SetRequestId(v string) *StartStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartStreamResponseBody) Validate() error {
	return dara.Validate(s)
}
