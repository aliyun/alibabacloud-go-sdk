// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetShareLinkTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExpiresIn(v int64) *GetShareLinkTokenResponseBody
	GetExpiresIn() *int64
	SetShareToken(v string) *GetShareLinkTokenResponseBody
	GetShareToken() *string
}

type GetShareLinkTokenResponseBody struct {
	// The validity period of the token. Unit: seconds. For example, a value of 7200 indicates two hours.
	//
	// example:
	//
	// 7200
	ExpiresIn *int64 `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
	// The JSON Web Token (JWT).
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJjOWI3YTVhYTA0ZDE0YWUzODY3ZmRjODg2ZmEwMWRhNCIsImN1c3RvbUpzb24iOiJ7XCJjbGllbnRJZFwiOlwiMjVkelgzdmJZcWt0Vnh5WFwiLFwiZG9tYWluSWRcIjpcImJqMjlcIixcInNjb3BlXCI6W1wiRFJJVkUuQUxMXCIsXCJTSEFSRS5BTExcIixcIkZJTEUuQUxMXCIsXCJVU0VSLkFMTFwiLFwiVklFVy5BTExcIixcIlNUT1JBR0UuQUxMXCIsXCJTVE9SQUdFRklMRS5MSVNUXCIsXCJCQVRDSFwiLFwiT0FVVEguQUxMXCIsXCJJTUFHRS5BTExcIixcIklOVklURS5BTExcIixcIkFDQ09VTlQuQUxMXCJdLFwicm9sZVwiOlwidXNlclwiLFwicmVmXCI6XCJodHRwczovL3d3dy5hbGl5dW5kcml2ZS5jb20vXCIsXCJkZXZpY2VfaWRcIjpcImIyODIwNWU1YzU5NzRjY2JiODI3MDNiNjhkYjhjNDUxXCJ9IiwiZXhwIjoxNjQ4NjE0NDkzLCJpYXQiOjE2NDg2MDcyMzN9.d3HVLvv_LFw2QhPrhvjH_kICWQJX9sKKt7NjQEqI_xE2JO_b7D8rPsFTZz93PLvZ7MhCmudTjGImUpd-ehFnI4Go-1S7BGaKaHFILvP-sWy18Wpikowjxx9mSbzBM_cO6D1LI-kyYhXKWHgVdADfVIPniTDA7-ffhUpi7cAebEs
	ShareToken *string `json:"share_token,omitempty" xml:"share_token,omitempty"`
}

func (s GetShareLinkTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetShareLinkTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetShareLinkTokenResponseBody) GetExpiresIn() *int64 {
	return s.ExpiresIn
}

func (s *GetShareLinkTokenResponseBody) GetShareToken() *string {
	return s.ShareToken
}

func (s *GetShareLinkTokenResponseBody) SetExpiresIn(v int64) *GetShareLinkTokenResponseBody {
	s.ExpiresIn = &v
	return s
}

func (s *GetShareLinkTokenResponseBody) SetShareToken(v string) *GetShareLinkTokenResponseBody {
	s.ShareToken = &v
	return s
}

func (s *GetShareLinkTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
