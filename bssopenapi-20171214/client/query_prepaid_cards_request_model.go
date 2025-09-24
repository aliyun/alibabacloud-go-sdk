// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPrepaidCardsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEffectiveOrNot(v bool) *QueryPrepaidCardsRequest
	GetEffectiveOrNot() *bool
	SetExpiryTimeEnd(v string) *QueryPrepaidCardsRequest
	GetExpiryTimeEnd() *string
	SetExpiryTimeStart(v string) *QueryPrepaidCardsRequest
	GetExpiryTimeStart() *string
}

type QueryPrepaidCardsRequest struct {
	// Specifies whether the prepaid card takes effect. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EffectiveOrNot *bool `json:"EffectiveOrNot,omitempty" xml:"EffectiveOrNot,omitempty"`
	// The end of the expiration time of prepaid cards to query. The value must be in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2018-08-01T00:00:00Z
	ExpiryTimeEnd *string `json:"ExpiryTimeEnd,omitempty" xml:"ExpiryTimeEnd,omitempty"`
	// The start of the expiration time of prepaid cards to query. The value must be in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2018-08-01T00:00:00Z
	ExpiryTimeStart *string `json:"ExpiryTimeStart,omitempty" xml:"ExpiryTimeStart,omitempty"`
}

func (s QueryPrepaidCardsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPrepaidCardsRequest) GoString() string {
	return s.String()
}

func (s *QueryPrepaidCardsRequest) GetEffectiveOrNot() *bool {
	return s.EffectiveOrNot
}

func (s *QueryPrepaidCardsRequest) GetExpiryTimeEnd() *string {
	return s.ExpiryTimeEnd
}

func (s *QueryPrepaidCardsRequest) GetExpiryTimeStart() *string {
	return s.ExpiryTimeStart
}

func (s *QueryPrepaidCardsRequest) SetEffectiveOrNot(v bool) *QueryPrepaidCardsRequest {
	s.EffectiveOrNot = &v
	return s
}

func (s *QueryPrepaidCardsRequest) SetExpiryTimeEnd(v string) *QueryPrepaidCardsRequest {
	s.ExpiryTimeEnd = &v
	return s
}

func (s *QueryPrepaidCardsRequest) SetExpiryTimeStart(v string) *QueryPrepaidCardsRequest {
	s.ExpiryTimeStart = &v
	return s
}

func (s *QueryPrepaidCardsRequest) Validate() error {
	return dara.Validate(s)
}
