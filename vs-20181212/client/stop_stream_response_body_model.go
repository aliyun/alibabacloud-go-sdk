// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *StopStreamResponseBody
	GetId() *string
	SetRequestId(v string) *StopStreamResponseBody
	GetRequestId() *string
}

type StopStreamResponseBody struct {
	// example:
	//
	// 32388487****92997-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopStreamResponseBody) GoString() string {
	return s.String()
}

func (s *StopStreamResponseBody) GetId() *string {
	return s.Id
}

func (s *StopStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopStreamResponseBody) SetId(v string) *StopStreamResponseBody {
	s.Id = &v
	return s
}

func (s *StopStreamResponseBody) SetRequestId(v string) *StopStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopStreamResponseBody) Validate() error {
	return dara.Validate(s)
}
