// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSuppressionListLevelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSuppressionListLevelResponseBody
	GetRequestId() *string
	SetSuppressionListLevel(v string) *GetSuppressionListLevelResponseBody
	GetSuppressionListLevel() *string
}

type GetSuppressionListLevelResponseBody struct {
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuppressionListLevel *string `json:"SuppressionListLevel,omitempty" xml:"SuppressionListLevel,omitempty"`
}

func (s GetSuppressionListLevelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSuppressionListLevelResponseBody) GoString() string {
	return s.String()
}

func (s *GetSuppressionListLevelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSuppressionListLevelResponseBody) GetSuppressionListLevel() *string {
	return s.SuppressionListLevel
}

func (s *GetSuppressionListLevelResponseBody) SetRequestId(v string) *GetSuppressionListLevelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSuppressionListLevelResponseBody) SetSuppressionListLevel(v string) *GetSuppressionListLevelResponseBody {
	s.SuppressionListLevel = &v
	return s
}

func (s *GetSuppressionListLevelResponseBody) Validate() error {
	return dara.Validate(s)
}
