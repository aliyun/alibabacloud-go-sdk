// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMassPushV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessageIds(v []*string) *MassPushV2ResponseBody
	GetMessageIds() []*string
	SetRequestId(v string) *MassPushV2ResponseBody
	GetRequestId() *string
}

type MassPushV2ResponseBody struct {
	MessageIds []*string `json:"MessageIds,omitempty" xml:"MessageIds,omitempty" type:"Repeated"`
	RequestId  *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MassPushV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MassPushV2ResponseBody) GoString() string {
	return s.String()
}

func (s *MassPushV2ResponseBody) GetMessageIds() []*string {
	return s.MessageIds
}

func (s *MassPushV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MassPushV2ResponseBody) SetMessageIds(v []*string) *MassPushV2ResponseBody {
	s.MessageIds = v
	return s
}

func (s *MassPushV2ResponseBody) SetRequestId(v string) *MassPushV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *MassPushV2ResponseBody) Validate() error {
	return dara.Validate(s)
}
