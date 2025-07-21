// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAapName(v string) *ListClientKeysRequest
	GetAapName() *string
}

type ListClientKeysRequest struct {
	// The name of the application access point (AAP).
	//
	// example:
	//
	// aap_test
	AapName *string `json:"AapName,omitempty" xml:"AapName,omitempty"`
}

func (s ListClientKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClientKeysRequest) GoString() string {
	return s.String()
}

func (s *ListClientKeysRequest) GetAapName() *string {
	return s.AapName
}

func (s *ListClientKeysRequest) SetAapName(v string) *ListClientKeysRequest {
	s.AapName = &v
	return s
}

func (s *ListClientKeysRequest) Validate() error {
	return dara.Validate(s)
}
