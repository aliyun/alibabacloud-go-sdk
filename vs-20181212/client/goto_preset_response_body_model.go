// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGotoPresetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GotoPresetResponseBody
	GetId() *string
	SetRequestId(v string) *GotoPresetResponseBody
	GetRequestId() *string
}

type GotoPresetResponseBody struct {
	// example:
	//
	// 348*****380-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 30295DF1-1DC7-48BA-BE5A-D58E61EB2375
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GotoPresetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GotoPresetResponseBody) GoString() string {
	return s.String()
}

func (s *GotoPresetResponseBody) GetId() *string {
	return s.Id
}

func (s *GotoPresetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GotoPresetResponseBody) SetId(v string) *GotoPresetResponseBody {
	s.Id = &v
	return s
}

func (s *GotoPresetResponseBody) SetRequestId(v string) *GotoPresetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GotoPresetResponseBody) Validate() error {
	return dara.Validate(s)
}
