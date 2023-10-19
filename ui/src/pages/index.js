import React from "react"
import { API } from '@stoplight/elements';
require('@stoplight/elements/styles.min.css');


export default function Home() {
  return <div>
      <API 
        apiDescriptionUrl="/swagger.json"
        router={typeof window === 'undefined' ? 'memory' : 'history'} 
      />
    </div>
}
