// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPresetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *SetPresetResponseBody
	GetId() *string
	SetRequestId(v string) *SetPresetResponseBody
	GetRequestId() *string
}

type SetPresetResponseBody struct {
	// example:
	//
	// 348*****380-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 30295DF1-1DC7-48BA-BE5A-D58E61EB2375
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetPresetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetPresetResponseBody) GoString() string {
	return s.String()
}

func (s *SetPresetResponseBody) GetId() *string {
	return s.Id
}

func (s *SetPresetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetPresetResponseBody) SetId(v string) *SetPresetResponseBody {
	s.Id = &v
	return s
}

func (s *SetPresetResponseBody) SetRequestId(v string) *SetPresetResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetPresetResponseBody) Validate() error {
	return dara.Validate(s)
}
