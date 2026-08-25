// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteModelsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *BatchDeleteModelsShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *BatchDeleteModelsShrinkRequest
	GetClientToken() *string
}

type BatchDeleteModelsShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// client-token-1
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s BatchDeleteModelsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteModelsShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteModelsShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *BatchDeleteModelsShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *BatchDeleteModelsShrinkRequest) SetBodyShrink(v string) *BatchDeleteModelsShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *BatchDeleteModelsShrinkRequest) SetClientToken(v string) *BatchDeleteModelsShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *BatchDeleteModelsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
