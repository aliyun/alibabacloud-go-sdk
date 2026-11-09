// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ApplyTokenResponseBody
	GetRequestId() *string
	SetToken(v string) *ApplyTokenResponseBody
	GetToken() *string
}

type ApplyTokenResponseBody struct {
	// The common parameter. The ID of each request is unique.
	//
	// example:
	//
	// 31782AAF-D0CC-44C3-ABFD-1B500276****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The token returned by the broker.
	//
	// > Do not make any assumptions about the length, format, or rules of the returned token. The actual returned value prevails.
	//
	// example:
	//
	// LzMT+XLFl5s/YWJ/MlDz4t/Lq5HC1iGU1P28HAMaxYxn8aQbALNtml7QZKl9L9kPe6LqUb95tEVo+zUqOogs9+jZwDUSzsd4X4qaD3n2TrBEuMOqKkk1Xdrvu9VBQQvIYbz7MJWZDYC3DlW7gLEr33Cuj54iIhagtBi3epStJitsssWs7otY9zhKOSZxhr49G3d0bh35mwyP18EMvDas8UlzeSozsSrujNUqZXOGK0PEBSd+rWMGDJlCt6GFmJgm2JFY7PJwf/7OOSmUYIYFs5o/PuPpoTMF+hcVXMs+0yDukIMTOzG9m3t8k36PVrghFmnK6pC3Rt3mibjW****ng==
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s ApplyTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyTokenResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyTokenResponseBody) GetToken() *string {
	return s.Token
}

func (s *ApplyTokenResponseBody) SetRequestId(v string) *ApplyTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyTokenResponseBody) SetToken(v string) *ApplyTokenResponseBody {
	s.Token = &v
	return s
}

func (s *ApplyTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
