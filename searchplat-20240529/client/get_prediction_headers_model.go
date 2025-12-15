// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPredictionHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetPredictionHeaders
	GetCommonHeaders() map[string]*string
	SetToken(v string) *GetPredictionHeaders
	GetToken() *string
}

type GetPredictionHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Token         *string            `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetPredictionHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetPredictionHeaders) GoString() string {
	return s.String()
}

func (s *GetPredictionHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetPredictionHeaders) GetToken() *string {
	return s.Token
}

func (s *GetPredictionHeaders) SetCommonHeaders(v map[string]*string) *GetPredictionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPredictionHeaders) SetToken(v string) *GetPredictionHeaders {
	s.Token = &v
	return s
}

func (s *GetPredictionHeaders) Validate() error {
	return dara.Validate(s)
}
