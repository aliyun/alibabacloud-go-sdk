// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSuppressionListLevelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetSuppressionListLevelResponseBody
	GetRequestId() *string
	SetSuppressionListLevel(v string) *SetSuppressionListLevelResponseBody
	GetSuppressionListLevel() *string
}

type SetSuppressionListLevelResponseBody struct {
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuppressionListLevel *string `json:"SuppressionListLevel,omitempty" xml:"SuppressionListLevel,omitempty"`
}

func (s SetSuppressionListLevelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetSuppressionListLevelResponseBody) GoString() string {
	return s.String()
}

func (s *SetSuppressionListLevelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetSuppressionListLevelResponseBody) GetSuppressionListLevel() *string {
	return s.SuppressionListLevel
}

func (s *SetSuppressionListLevelResponseBody) SetRequestId(v string) *SetSuppressionListLevelResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetSuppressionListLevelResponseBody) SetSuppressionListLevel(v string) *SetSuppressionListLevelResponseBody {
	s.SuppressionListLevel = &v
	return s
}

func (s *SetSuppressionListLevelResponseBody) Validate() error {
	return dara.Validate(s)
}
